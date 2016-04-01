package util

import (
	"encoding/binary"
	"math"
)

func Bytes4ToFloat32(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
func Float32Bytes4(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, bits)
	return bytes
}
func Bytes8ToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
func Float64ToBytes8(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}
