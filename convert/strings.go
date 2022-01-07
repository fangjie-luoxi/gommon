package convert

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"math"
	"strings"
	"unsafe"
)

// MD5Hash md5加密
func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// BytesToStr 将字节切片转换为字符串而不分配内存.
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StrToBytes 在没有内存分配的情况下将字符串转换为字节切片.
func StrToBytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// SnakeCasedName 将 String 转换为 小驼峰
// ex: FooBar -> foo_bar
func SnakeCasedName(name string) string {
	newstr := make([]byte, 0, len(name)+1)
	for i := 0; i < len(name); i++ {
		c := name[i]
		if isUpper := 'A' <= c && c <= 'Z'; isUpper {
			if i > 0 {
				newstr = append(newstr, '_')
			}
			c += 'a' - 'A'
		}
		newstr = append(newstr, c)
	}

	return BytesToStr(newstr)
}

// TitleCasedName 将字符串转换为标题大小写
// ex: foo_bar -> FooBar
func TitleCasedName(name string) string {
	newstr := make([]byte, 0, len(name))
	upNextChar := true

	name = strings.ToLower(name)

	for i := 0; i < len(name); i++ {
		c := name[i]
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= c && c <= 'z' {
				c -= 'a' - 'A'
			}
		case c == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, c)
	}

	return BytesToStr(newstr)
}

// Float64ToByte 将 float64 转换为字节
// ref: https://stackoverflow.com/questions/43693360/convert-float64-to-byte-array
func Float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

// ByteToFloat64 将字节转换为 float64
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
