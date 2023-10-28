package utils

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"sort"
	"strconv"
)

var hexNumbers = []byte("0123456789ABCDEF")

// BToU16 从字节切片中获取uint16
func BToU16(b []byte, index int, offset int) uint16 {
	return binary.BigEndian.Uint16(b[index:offset])
}

// BToU32 从字节切片中获取uint32
func BToU32(b []byte, index int, offset int) uint32 {
	return binary.BigEndian.Uint32(b[index:offset])
}

// BToU64 从字节切片中获取uint64
func BToU64(b []byte, index int, offset int) uint64 {
	return binary.BigEndian.Uint64(b[index:offset])
}

// U16ToB 将uint16转换为字节切片
func U16ToB(v uint16, b []byte) {
	binary.BigEndian.PutUint16(b, v)
}

// U32ToB 将uint32转换为字节切片
func U32ToB(v uint32, b []byte) {
	binary.BigEndian.PutUint32(b, v)
}

// U64ToB 将uint64转换为字节切片
func U64ToB(v uint64, b []byte) {
	binary.BigEndian.PutUint64(b, v)
}

// StringToInt 字符串转int
func StringToInt(str string) int {
	var i int
	var err error
	if i, err = strconv.Atoi(str); err != nil {
		return 0
	}
	return i
}

// IntToBytes int转字节切片
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// Int64ToBytes int64转字节切片
func Int64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// ByteToHex 编码
func ByteToHex(value byte) []byte {
	buf := make([]byte, 2)
	buf[0] = hexNumbers[value>>4]
	buf[1] = hexNumbers[value&0x0F]
	return buf
}

// GbkToUtf8 gbk转utf8
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// InStringArray 判断字符串是否在数组中
func InStringArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

// HexToFloat64 16进制转float64
func HexToFloat64(hex []byte) float64 {
	var f float64
	_ = binary.Read(bytes.NewBuffer(hex), binary.BigEndian, &f)
	return f
}
