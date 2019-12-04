package utils

import (
	m "crypto/md5"
	"fmt"
)

//MD5 加密
func Md5(str string) string {
	//方法一
	data := []byte(str)
	has := m.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
