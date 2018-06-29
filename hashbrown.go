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
		// a, b, c, and d are the output buffers
		// k is which word is being used from the message block
		// s is the amount to rotate by
		// i is the value from the T table
		// X is the message block
		// T is a table with MAGIC values
		// op is the function to call on the buffers (one of F, G, H, or I)
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

	wordByteSize := 4   // number of bytes in 32-bit word
	blockWordSize := 16 // number of  32-bit words in a block
	// It's convenient to have padding length in bytes
	lengthInBits := uint64(len(message) * 8)

	// Here's how padding works:
	// the message is padded (extended) so that the length of the message in bits
	// modulo 512 is 448. In other words len_bits(message) % 512 == 448 after padding
	// 512 bits is 64 bytes, 448 bits is 56 bytes
	message = append(message, 0x80)
	for len(message)%64 != 56 {
		message = append(message, 0)
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

	rotateAmounts := [64]int{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}
	// for each 16 word block in the message
	for i := 0; i < N/blockWordSize; i++ {
		// Get a slice of the block. Go slices are just views into the underlying
		// array storage, and thus this saves us a copy.
		blockAddress := i * blockWordSize * wordByteSize
		block := message[blockAddress : blockAddress+blockWordSize*wordByteSize]
		// Turn the block from a slice of bytes into a slice of words
		X := make([]uint32, blockWordSize)
		binary.Read(bytes.NewReader(block), binary.LittleEndian, &X)
		AA := A
		BB := B
		CC := C
		DD := D
		// ready to begin implementing rounds
		// how should I for loop this

		for k := 0; k < 64; k++ {
			var g int
			var op MD5Operation
			if 0 <= k && k <= 15 {
				g = k
				op = F
			} else if 16 <= k && k <= 31 {
				g = (5*k + 1) % 16
				op = G
			} else if 32 <= k && k <= 47 {
				g = (3*k + 5) % 16
				op = H
			} else if 48 <= k && k <= 63 {
				g = (7 * k) % 16
				op = I
			}

			res := Op(AA, BB, CC, DD, g, rotateAmounts[k], k, X, T, op)
			AA = DD
			DD = CC
			CC = BB
			BB = res

		}

		A += AA
		B += BB
		C += CC
		D += DD
	}
	digest := make([]byte, 16)
	binary.LittleEndian.PutUint32(digest, A)
	binary.LittleEndian.PutUint32(digest, B)
	binary.LittleEndian.PutUint32(digest, C)
	binary.LittleEndian.PutUint32(digest, D)
	return digest
}

func main() {
	return
}
