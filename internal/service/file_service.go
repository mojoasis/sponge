package service

import (
	"context"
	"sponge/internal/model/dto"
	"sponge/pkg/global"
	"sponge/pkg/utils"

	"go.uber.org/zap"
)

// UploadVideo 上传视频
func UploadVideo(ctx context.Context, uid int64, req *dto.UploadFileReq) []dto.UploadVideoRes {
	var results []dto.UploadVideoRes
	// 循环处理上传的文件
	for _, file := range req.Files {
		res, err := utils.UploadAndGenCover(ctx, uid, file)
		if err != nil {
			global.Logger.Error("视频处理流程异常",
				zap.Int64("uid", uid),
				zap.String("filename", file.Filename),
				zap.Error(err),
			)
			continue
		}
		results = append(results, res)
	}
	return results
}
