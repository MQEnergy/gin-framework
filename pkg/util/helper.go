package util

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
