package main

import "testing"
import "bytes"

// TestMD5Sum tests that MD5Sum works correctly
func TestMD5Sum(t *testing.T) {
	cases := [][][]byte{
		{[]byte(""), []byte("d41d8cd98f00b204e9800998ecf8427e")},
		{[]byte("a"), []byte("0cc175b9c0f1b6a831c399e269772661")},
		{[]byte("abc"), []byte("900150983cd24fb0d6963f7d28e17f72")},
		{[]byte("message digest"), []byte("f96b697d7cb7938d525a2f31aaf161d0")},
		{[]byte("abcdefghijklmnopqrstuvwxyz"), []byte("c3fcd3d76192e4007dfb496cca67e13b")},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), []byte("d174ab98d277d9f5a5611c2c9f419d9f")},
		{[]byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890"), []byte("57edf4a22be3c955ac49da2e2107b67a")},
		{[]byte("U5t0joJJiKSJxSnG2DnyVyQ1BUIJTCbgfvWmxJTZMBe8hDvCdIxkkhgK7du8ux8DerIN6SoIFlf0S3C83iEndzAYfyIq3ZGI9U12NAKdfiRHKH8RbzITvymqO9VUrkcHPPnPZ9GMctKJzv7ejZGPJ9XqRWwEBSznv4cKZluK37w7DTxJExqASxs1SMHwPnbDBibPhhlz"), []byte("6F43F82BA545066274C2E61BACD6D39B")},
	}
	for i := 0; i < len(cases); i++ {
		input := cases[i][0]
		expected := cases[i][1]
		sum := MD5Sum(input)
		if bytes.Equal(sum, expected) {
			t.Fail()
		}
	}
}
