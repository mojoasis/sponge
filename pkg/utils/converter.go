package utils

import "github.com/jinzhu/copier"

// ToDTO 将 Model (数据库实体) 转换为 DTO (展示对象)
func ToDTO(src interface{}, dst interface{}) error {
	return copier.Copy(dst, src)
}

// ToModel 将 DTO (请求对象) 转换为 Model (数据库实体)
func ToModel(src interface{}, dst interface{}) error {
	return copier.Copy(dst, src)
}

// ToList 批量转换工具 (比如 []Model 转 []DTO)
func ToList(src interface{}, dst interface{}) error {
	return copier.Copy(dst, src)
}

// CopyToModelForUpdate 专门用于更新场景
// 它会跳过 src (DTO) 中的零值字段，只将有值的字段覆盖到 dst (Model)
func CopyToModelForUpdate(src interface{}, dst interface{}) error {
	return copier.CopyWithOption(dst, src, copier.Option{
		IgnoreEmpty: true, // 核心：忽略 DTO 中的空值
		DeepCopy:    true,
	})
}
