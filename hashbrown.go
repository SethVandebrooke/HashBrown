package main

import "encoding/binary"

// MD5Sum computes the MD5 hash of the input
func MD5Sum(message []byte) []byte {

	bufferA := [4]uint32{0x01, 0x23, 0x45, 0x67}
	bufferB := [4]uint32{0x89, 0xab, 0xcd, 0xef}
	bufferC := [4]uint32{0xfe, 0xdc, 0xba, 0x98}
	bufferD := [4]uint32{0x76, 0x54, 0x32, 0x10}

	lengthInBits := int64(len(message) * 8) // b from rfc1321
	paddingMult := lengthInBits / 512
	// It's convenient to have padding length in bytes
	paddingLength := (lengthInBits - ((paddingMult + 1) * 512) - 64) / 8
	for index := int64(0); index < paddingLength; index++ {
		paddingByte := byte(0)
		if index == 0 {
			mask := (byte)(1 << 0)
			paddingByte |= mask
		}
		message = append(message, paddingByte)
	}

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(lengthInBits))
	message = append(message, b...)

	return message
}

func main() {
	return
}
