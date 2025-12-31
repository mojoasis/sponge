package utils

import "golang.org/x/crypto/bcrypt"

// CryptPassword 加密密码
func CryptPassword(password string) (string, error) {
	cost := 10 // 生产环境建议 10-12
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPassword 校验密码
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
