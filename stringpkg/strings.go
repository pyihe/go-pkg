package stringpkg

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// Bytes string转换为[]byte
func Bytes(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// Uint64 string转换为uint64
func Uint64(s string) (uint64, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return uint64(val), nil
}

// Int64 string转换为int64
func Int64(s string) (int64, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

// Int string转换为int
func Int(s string) (int, error) {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return val, err
}

// Bool string转换为bool
func Bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// Float32 string转换为float32
func Float32(s string) (float32, error) {
	fv, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}
	return float32(fv), nil
}

// Float64 string转换为float64
func Float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Last 返回s中最后一个b的位置
func Last(s string, b byte) int {
	i := len(s)
	for i--; i >= 0; i-- {
		if s[i] == b {
			break
		}
	}
	return i
}

// IsEmpty 判断字符串本身或者去除空格后是否为空字符串
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	s = strings.TrimSpace(s)
	return s == ""
}

// SliceEqual 判断两个字符串切片值是否相等
func SliceEqual(a, b []string) bool {
	aLen, bLen := len(a), len(b)
	if aLen != bLen {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:aLen]
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// Contain 判断字符串切片是否包含ele
func Contain(ele string, s []string) bool {
	for _, v := range s {
		if v == ele {
			return true
		}
	}
	return false
}

// Reverse 字符切片反转
func Reverse(s []string) {
	l := len(s)
	for i := l/2 - 1; i >= 0; i-- {
		opp := l - i - 1
		s[i], s[opp] = s[opp], s[i]
	}
}

// Remove 从s中删除第一个ele
func Remove(ele string, s []string) {
	for i := range s {
		if s[i] == ele {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
}

// MaxSubStrLen 寻找最长不重复子串长度
func MaxSubStrLen(str string) int {
	var maxLen int
	var start int
	m := make(map[rune]int)
	for i, ch := range []rune(str) {
		if lastIndex, ok := m[ch]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		m[ch] = i
	}
	return maxLen
}

// FilterRepeatBySlice 元素去重
func FilterRepeatBySlice(slc []string) []string {
	var result []string
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

// FilterRepeatByMap 通过map元素去重
func FilterRepeatByMap(slc []string) []string {
	var result []string
	tempMap := make(map[string]byte, 0)
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
