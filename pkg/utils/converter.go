package utils

import (
	"github.com/jinzhu/copier"
)

// ToDTO 将 model 转换为 DTO
// 使用 copier 库进行结构体之间的字段复制
// 支持相同字段名的自动转换，忽略不存在的字段
func ToDTO(src interface{}, dst interface{}) error {
	return copier.Copy(dst, src)
}

// ToDTOWithIgnore 将 model 转换为 DTO，并忽略指定字段
func ToDTOWithIgnore(src interface{}, dst interface{}, ignoreFields ...string) error {
	opt := copier.Option{
		IgnoreEmpty: false,
		DeepCopy:    false,
	}
	return copier.CopyWithOption(dst, src, opt)
}
