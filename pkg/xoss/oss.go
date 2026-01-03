package xoss

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type OSSClient struct {
	S3Client      *s3.Client
	PresignClient *s3.PresignClient
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
		// 如果没有 http 开头，自动补齐
		o.BaseEndpoint = aws.String(endpoint) // 直接设置基础地址
		o.UsePathStyle = true                 // 关键：MinIO 等私有云通常需要开启路径风格访问
	})
	return client
}

// MakeBucket 创建桶
func (o *OSSClient) MakeBucket(ctx context.Context, bucketName string) {
	_, err := o.S3Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		_, err = o.S3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
		})
		if err != nil {
			log.Printf("创建存储桶 %s 失败: %v\n", bucketName, err)
			return
		}
		fmt.Printf("成功创建存储桶: %v\n", bucketName)
	}
}

// UploadFile 上传普通文件
func (o *OSSClient) UploadFile(ctx context.Context, bucketName string, objectKey string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = o.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	return err
}

// UploadStream 直接接收 io.Reader
func (o *OSSClient) UploadStream(ctx context.Context, bucketName string, objectKey string, stream io.Reader, size int64) error {
	_, err := o.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(objectKey),
		Body:          stream,
		ContentLength: aws.Int64(size), // 提供大小有助于 SDK 优化分片
	})
	return err
}

// GetPresignedURL 生成预签名 URL 的方法
func (o *OSSClient) GetPresignedURL(ctx context.Context, bucketName string, objectKey string, expires time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(o.S3Client)
	request, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, s3.WithPresignExpires(expires))

	if err != nil {
		return "", err
	}
	return request.URL, nil
}

// ConcurrentUpload 并发上传
func (o *OSSClient) ConcurrentUpload(ctx context.Context, bucketName string, key string, file io.Reader) error {
	uploader := manager.NewUploader(o.S3Client, func(u *manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024
		u.Concurrency = 5
	})

	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

// GenCover 使用 FFmpeg 截取视频第一帧并上传
//func (o *OSSClient) GenCover(ctx context.Context, videoPath string, snapshotPath string) (string, error) {
//	// 1. 抽帧保存到本地临时文件
//	err := ffmpeg.Input(videoPath).
//		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 0)}).
//		Output(snapshotPath, ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
//		OverWriteOutput().Run()
//	if err != nil {
//		return "", err
//	}
//
//	// 2. 上传到 OSS
//	objectKey := "snapshots/" + snapshotPath
//	err = o.UploadFile(ctx, objectKey, "portal", snapshotPath)
//	return objectKey, err
//}
