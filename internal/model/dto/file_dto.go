package dto

import "mime/multipart"

// UploadFileReq 上传多文件请求
type UploadFileReq struct {
	// 使用 FileHeader 指针的切片
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

// UploadBaseRes 上传基本文件（图片、文档等）
type UploadBaseRes struct {
	FileBaseInfo // 匿名嵌套，JSON 会自动展开
}

// UploadVideoRes 视频类文件
type UploadVideoRes struct {
	FileBaseInfo
	CoverUrl string  `json:"coverUrl"`
	Duration float32 `json:"duration"`
}

// FileBaseInfo 基础文件属性
type FileBaseInfo struct {
	FileUrl string `json:"playUrl"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Size    int64  `json:"size"` // 建议增加文件大小，方便前端显示
}
