package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func GenPasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

// 密码验证 检查提供的密码是否与哈希密码匹配
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
