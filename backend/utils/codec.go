package utils

import "encoding/binary"

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
