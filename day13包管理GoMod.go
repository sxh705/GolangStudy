package main

import (
	"encoding/base64"
	"fmt"

	"github.com/forgoer/openssl"
)

/*
包管理工具提供对外部包的依赖管理和版本控制
从go1.13默认开启gomod
使用go env查看是否开启 go111moode=on
go mod init 模块名
*/

func main_day13() {
	src := []byte("123456")
	key := []byte("1234567890123456")
	dst, _ := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	fmt.Printf(base64.StdEncoding.EncodeToString(dst)) // yXVUkR45PFz0UfpbDB8/ew==

	dst, _ = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	fmt.Println(string(dst)) // 123456
}
