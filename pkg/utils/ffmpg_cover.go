package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sponge/internal/model/dto"
	"sponge/pkg/global"
	"sync"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

func UploadAndGenCover(ctx context.Context, userID int64, header *multipart.FileHeader) (res dto.UploadVideoRes, err error) {
	// 1. 创建临时视频文件路径
	// 使用 NextID 确保并发上传时文件名不冲突
	nextID := NextID()
	tmpVideoPath := fmt.Sprintf("tmp/v_%d.mp4", nextID)
	tmpCoverPath := fmt.Sprintf("tmp/c_%d.jpg", nextID)

	// 确保函数退出时删除临时文件
	defer os.Remove(tmpVideoPath)

	// 2. 将上传的文件保存到本地临时路径
	file, err := header.Open()
	if err != nil {
		return res, err
	}
	defer file.Close()

	out, err := os.Create(tmpVideoPath)
	if err != nil {
		return res, fmt.Errorf("创建临时文件失败: %v", err)
	}
	_, err = io.Copy(out, file)
	out.Close() // 必须在处理前关闭，确保数据完全写入磁盘
	if err != nil {
		return res, fmt.Errorf("保存临时文件失败: %v", err)
	}

	// 3. 提取元数据 (宽高和时长)
	metadata, err := GetFileMeta(tmpVideoPath)
	if err != nil {
		global.Logger.Warn("无法获取文件元数据", zap.Error(err))
		metadata = &FileMetadata{} // 防止空指针，使用零值
	}

	// 4. 定义 OSS 存储路径
	videoKey := fmt.Sprintf("videos/%d/%d_%s", userID, nextID, header.Filename)
	coverKey := fmt.Sprintf("covers/%d/%d.jpg", userID, nextID)

	// 5. 并发执行：上传 OSS 和 FFmpeg 抽帧
	var wg sync.WaitGroup
	wg.Add(2)
	errChan := make(chan error, 2)

	// 协程 A：上传视频文件到 OSS
	go func() {
		defer wg.Done()
		vFile, err := os.Open(tmpVideoPath)
		if err != nil {
			errChan <- err
			return
		}
		defer vFile.Close()

		if err := global.OSS.ConcurrentUpload(ctx, global.Config.OSS.VideoBucket, videoKey, vFile); err != nil {
			errChan <- fmt.Errorf("视频上传失败: %v", err)
		}
	}()

	// 协程 B：FFmpeg 抽帧并上传封面
	go func() {
		defer wg.Done()
		err := ffmpeg.Input(tmpVideoPath).
			Output(tmpCoverPath, ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
			Run()

		if err != nil {
			errChan <- fmt.Errorf("FFmpeg 抽帧失败: %v", err)
			return
		}
		defer os.Remove(tmpCoverPath)

		if err := global.OSS.UploadFile(ctx, global.Config.OSS.ImgBucket, coverKey, tmpCoverPath); err != nil {
			errChan <- fmt.Errorf("封面上传失败: %v", err)
		}
	}()

	wg.Wait()
	close(errChan)

	// 6. 检查是否有错误发生
	for e := range errChan {
		if e != nil {
			return res, e
		}
	}

	// 7. 组装返回结果
	res = dto.UploadVideoRes{
		FileBaseInfo: dto.FileBaseInfo{
			FileUrl: videoKey,
			Width:   metadata.Width,
			Height:  metadata.Height,
		},
		CoverUrl: coverKey,
		Duration: metadata.Duration,
	}
	return res, nil
}

// UploadAndGenCover 流式处理视频：上传的同时抽帧
// userID 从 API 层通过 res.GetUserID(c) 获取后传入
//func UploadAndGenCover(ctx context.Context, userID int64, header *multipart.FileHeader) (videoKey, coverKey string, err error) {
//	file, err := header.Open()
//	if err != nil {
//		return "", "", err
//	}
//	defer file.Close()
//
//	// 1. 定义存储路径 (根据 userID 分文件夹存储)
//	videoKey = fmt.Sprintf("videos/%d/%d_%s", userID, NextID(), header.Filename)
//	coverKey = fmt.Sprintf("covers/%d/%d.jpg", userID, NextID())
//
//	// 2. 创建管道
//	// pr: 给 FFmpeg 读取
//	// pw: 给 TeeReader 写入
//	pr, pw := io.Pipe()
//	// TeeReader：读取 file 的同时，数据自动流向 pw
//	tr := io.TeeReader(file, pw)
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	// 用于收集并发中的错误
//	errChan := make(chan error, 2)
//
//	// 协程 A：执行 OSS 分片流式上传 (数据源是 tr)
//	go func() {
//		defer wg.Done()
//		defer pw.Close() // 上传完成或失败后关闭管道，通知 FFmpeg 结束读取
//		if err := global.OSS.ConcurrentUpload(ctx, global.Config.OSS.VideoBucket, videoKey, tr); err != nil {
//			errChan <- fmt.Errorf("视频上传失败: %v", err)
//		}
//	}()
//
//	// 协程 B：执行 FFmpeg 实时抽帧 (数据源是 pr)
//	go func() {
//		defer wg.Done()
//		tmpCoverPath := fmt.Sprintf("tmp/%d.jpg", NextID())
//		// FFmpeg 从标准输入 (pipe:0) 读取流数据
//		err := ffmpeg.Input("pipe:0").
//			Output(tmpCoverPath, ffmpeg.KwArgs{
//				"vframes": 1,
//				"format":  "image2",
//				"vcodec":  "mjpeg",
//			}).
//			WithInput(pr).
//			Run()
//
//		if err != nil {
//			errChan <- fmt.Errorf("FFmpeg 抽帧失败: %v", err)
//			return
//		}
//		defer os.Remove(tmpCoverPath)
//
//		// 将抽好的本地图片上传到 OSS
//		if err := global.OSS.UploadFile(ctx, global.Config.OSS.ImgBucket, coverKey, tmpCoverPath); err != nil {
//			errChan <- fmt.Errorf("封面上传失败: %v", err)
//		}
//	}()
//
//	// 等待两个协程完成
//	wg.Wait()
//	close(errChan)
//
//	// 检查是否有错误发生
//	for e := range errChan {
//		if e != nil {
//			return "", "", e
//		}
//	}
//
//	return videoKey, coverKey, nil
//}
