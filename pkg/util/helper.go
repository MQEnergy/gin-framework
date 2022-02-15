package util

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hashicorp/go-uuid"
	"io"
)

// InStringSlice 判断某个字符串是否在字符串切片中
func InStringSlice(needle string, heyhack []string) bool {
	for _, v := range heyhack {
		if v == needle {
			return true
		}
	}
	return false
}

// InIntSlice 判断某个数值是否在整型切片中
func InIntSlice(needle int, heyhack []int) bool {
	for _, v := range heyhack {
		if v == needle {
			return true
		}
	}
	return false
}

// GenerateBaseSnowId 生成雪花算法ID
func GenerateBaseSnowId(num int) string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	id := node.Generate()
	switch num {
	case 2:
		return id.Base2()
	case 32:
		return id.Base32()
	case 36:
		return id.Base36()
	case 58:
		return id.Base58()
	case 64:
		return id.Base64()
	default:
		return id.Base32()
	}
}

// GenerateUuid 生成随机字符串
func GenerateUuid(size int) string {
	str, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return gstr.SubStr(str, 0, size)
}

// GeneratePasswordHash 生成密码hash值
func GeneratePasswordHash(password string, salt string) string {
	md5 := md5.New()
	io.WriteString(md5, password)
	str := fmt.Sprintf("%x", md5.Sum(nil))
	s := sha256.New()
	io.WriteString(s, password+salt)
	str = fmt.Sprintf("%x", s.Sum(nil))
	return str
}
