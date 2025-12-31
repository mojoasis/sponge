package xoss

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type OSSClient struct {
	S3Client *s3.Client
	Bucket   string
}

// NewOSSClient 使用 BaseEndpoint 方式初始化
func NewOSSClient(endpoint, ak, sk string) *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// 1. 设置凭证
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(ak, sk, "")),
		// 2. 设置区域（S3 协议必须有，私有云一般填 us-east-1 或 auto）
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil
	}

	// 3. 在创建 S3 客户端时直接指定 BaseEndpoint
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint) // 直接设置基础地址
		o.UsePathStyle = true                 // 关键：MinIO 等私有云通常需要开启路径风格访问
	})

	return client
}

// UploadFile 上传普通文件
func (o *OSSClient) UploadFile(ctx context.Context, objectKey string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = o.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(o.Bucket),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	return err
}

// GenCover 使用 FFmpeg 截取视频第一帧并上传
func (o *OSSClient) GenCover(ctx context.Context, videoPath string, snapshotPath string) (string, error) {
	// 1. 抽帧保存到本地临时文件
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 0)}).
		Output(snapshotPath, ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		OverWriteOutput().Run()
	if err != nil {
		return "", err
	}

	// 2. 上传到 OSS
	objectKey := "snapshots/" + snapshotPath
	err = o.UploadFile(ctx, objectKey, snapshotPath)
	return objectKey, err
}
