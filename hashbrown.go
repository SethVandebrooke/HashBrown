package main

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/bits"
)

type MD5Operation = func(X, Y, Z uint32) uint32

// MD5Sum computes the MD5 hash of the input
func MD5Sum(message []byte) []byte {

	var A uint32 = 0x76543210
	var B uint32 = 0xfecdab98
	var C uint32 = 0x89abcdef
	var D uint32 = 0x01234567

	F := func(X, Y, Z uint32) uint32 {
		return (X & Y) | (^X)&Z
	}

	Op := func(a, b, c, d uint32, k, s, i int, X, T []uint32, op MD5Operation) uint32 {
		return b + bits.RotateLeft32((a+op(b, c, d)+X[k]+T[i]), s)
	}

	G := func(X, Y, Z uint32) uint32 {
		return (X & Z) | Y&(^Z)
	}

	H := func(X, Y, Z uint32) uint32 {
		return X ^ Y ^ Z
	}

	I := func(X, Y, Z uint32) uint32 {
		return Y ^ (X | (^Z))
	}

	lengthInBits := uint64(len(message) * 8) // b from rfc1321
	paddingMult := lengthInBits / 512
	wordByteSize := 4   // number of bytes in 32-bit word
	blockWordSize := 16 // number of  32-bit words in a block
	// It's convenient to have padding length in bytes
	paddingLength := (lengthInBits - ((paddingMult + 1) * 512) - 64) / 8
	for index := uint64(0); index < paddingLength; index++ {
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
	// N is the number of 32-bit words in the message,
	N := ((len(message) * 8) / 32)
	T := make([]uint32, 64)
	for i := 0; i < len(T); i++ {
		// Type conversion. Type conversion to the max!
		T[i] = uint32(math.Floor(math.Pow(2, 32) * math.Abs(math.Sin(float64(i+1)))))
	}

	// for each 16 word block in the message (should be 512 bits)
	for i := 0; i < N/blockWordSize; i++ {
		// Get a slice of the block. Go slices are just views into the underlying
		// array storage, and thus this saves us a copy.
		blockAddress := i * blockWordSize * wordByteSize
		block := message[blockAddress : blockAddress+blockWordSize*wordByteSize]
		// Turn the block from an slice of bytes into a slice of words
		X := make([]uint32, blockWordSize)
		binary.Read(bytes.NewReader(block), binary.LittleEndian, &X)
		AA := A
		BB := B
		CC := C
		DD := D

		// ready to begin implementing rounds
	}

	return message
}

func main() {
	return
}
