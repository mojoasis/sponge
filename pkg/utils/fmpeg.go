package utils

import (
	"fmt"
	"strconv"

	"github.com/bytedance/sonic"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// FileMetadata 定义提取出的元数据
type FileMetadata struct {
	Width    int
	Height   int
	Duration float32 // 秒
}

// GetFileMeta 使用 ffprobe 提取视频或图片的元数据
func GetFileMeta(filePath string) (*FileMetadata, error) {
	dataStr, err := ffmpeg.Probe(filePath)
	if err != nil {
		return nil, fmt.Errorf("ffprobe 解析失败: %v", err)
	}

	data := []byte(dataStr)
	metadata := &FileMetadata{}
	root, err := sonic.Get(data)
	if err != nil {
		return nil, err
	}

	// 1. 获取宽高：直接通过路径获取第一个 stream 的数据
	// GetByPath 如果路径不存在，返回的 Node 调用 Int64() 会返回 0，不会 panic
	w, _ := root.GetByPath("streams", 0, "width").Int64()
	h, _ := root.GetByPath("streams", 0, "height").Int64()
	metadata.Width = int(w)
	metadata.Height = int(h)

	// 2. 获取时长
	durationNode := root.GetByPath("format", "duration")
	if durationNode.Valid() {
		dStr, _ := durationNode.String()
		d, _ := strconv.ParseFloat(dStr, 32)
		metadata.Duration = float32(d)
	}

	return metadata, nil
}
