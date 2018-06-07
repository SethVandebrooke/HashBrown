package main

import "testing"

// TestMD5Sum tests that MD5Sum works correctly
func TestMD5Sum(t *testing.T) {
	input := "U5t0joJJiKSJxSnG2DnyVyQ1BUIJTCbgfvWmxJTZMBe8hDvCdIxkkhgK7du8ux8DerIN6SoIFlf0S3C83iEndzAYfyIq3ZGI9U12NAKdfiRHKH8RbzITvymqO9VUrkcHPPnPZ9GMctKJzv7ejZGPJ9XqRWwEBSznv4cKZluK37w7DTxJExqASxs1SMHwPnbDBibPhhlz"
	expected := "6F43F82BA545066274C2E61BACD6D39B"
	sum := MD5Sum(input)
	if sum != expected {
		t.Fail()
	}
}
