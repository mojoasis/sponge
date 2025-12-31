package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// InitSnowflake 初始化雪花算法节点
func InitSnowflake(machineID int64) error {
	var err error
	// snowflake.NewNode 接收 0-1023 之间的 ID
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		return fmt.Errorf("failed to init snowflake: %w", err)
	}
	return nil
}

// NextID 生成 64 位整型 ID
func NextID() int64 {
	return node.Generate().Int64()
}

// NextIDStr 生成字符串格式 ID (常用于前端，防止精度丢失)
func NextIDStr() string {
	return node.Generate().String()
}
