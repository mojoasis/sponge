package model

import (
	"time"

	"sponge/pkg/utils"

	"gorm.io/gorm"
)

// BaseModel 基础模型，包含 ID、时间戳和软删除
type BaseModel struct {
	// ID 使用雪花算法生成，json 序列化为字符串防止前端精度丢失
	ID        int64          `gorm:"primaryKey;autoIncrement:false" json:"id,string"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除字段，json 不返回
}

// BeforeCreate GORM 钩子：在创建记录前自动填充雪花 ID
// 注意：必须使用指针接收者，否则无法回写 ID
func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// 1. 自动生成雪花 ID
	if m.ID == 0 {
		m.ID = utils.NextID()
	}

	// 2. 如果需要强制处理时间，可以在这里手动赋值（通常 GORM 会自动处理 CreatedAt）
	// m.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate GORM 钩子：更新前的逻辑
func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	// 预留给需要的业务逻辑，比如记录修改日志等
	return nil
}
