package util

import (
	"fmt"
	"testing"
)

func TestComparePassword(t *testing.T) {
	pwd := "123456"
	hashedPwd := GenPasswordHash(pwd)
	fmt.Println(hashedPwd)
	fmt.Println(ComparePassword(hashedPwd, pwd))
}
