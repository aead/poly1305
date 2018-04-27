// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package poly1305

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func fromHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

var mult = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
var sNul = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var sSet = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

var TagZeroByteMessage_1 = []byte{0x0f, 0x01, 0x23, 0x45, 0x07, 0x88, 0xab, 0xcd, 0x0f, 0x00, 0x23, 0x45, 0x07, 0x88, 0xab, 0xcd}
var TagZeroByteMessage_2 = []byte{0x0e, 0x01, 0x23, 0x45, 0x07, 0x88, 0xab, 0xcd, 0x0f, 0x00, 0x23, 0x45, 0x07, 0x88, 0xab, 0xcd}
var TagFirstBitMessage_1 = []byte{0x10, 0x24, 0x68, 0x4c, 0x8f, 0x33, 0x79, 0xdd, 0x0f, 0x23, 0x68, 0x4c, 0x8f, 0x33, 0x79, 0xdd}
var TagFirstBitMessage_2 = []byte{0x0f, 0x24, 0x68, 0x4c, 0x8f, 0x33, 0x79, 0xdd, 0x0f, 0x23, 0x68, 0x4c, 0x8f, 0x33, 0x79, 0xdd}
var TagSecndBitMessage_1 = []byte{0x11, 0x47, 0xad, 0x53, 0x17, 0xdf, 0x46, 0xed, 0x0f, 0x46, 0xad, 0x53, 0x17, 0xdf, 0x46, 0xed}
var TagSecndBitMessage_2 = []byte{0x10, 0x47, 0xad, 0x53, 0x17, 0xdf, 0x46, 0xed, 0x0f, 0x46, 0xad, 0x53, 0x17, 0xdf, 0x46, 0xed}

var vectors = []struct {
	msg, key, tag []byte
}{
	{
		[]byte("Hello world!"),
		[]byte("this is 32-byte key for Poly1305"),
		[]byte{0xa6, 0xf7, 0x45, 0x00, 0x8f, 0x81, 0xc9, 0x16, 0xa2, 0x0d, 0xcc, 0x74, 0xee, 0xf2, 0xb2, 0xf0},
	},
	{
		make([]byte, 32),
		[]byte("this is 32-byte key for Poly1305"),
		[]byte{0x49, 0xec, 0x78, 0x09, 0x0e, 0x48, 0x1e, 0xc6, 0xc2, 0x6b, 0x33, 0xb9, 0x1c, 0xcc, 0x03, 0x07},
	},
	{
		make([]byte, 2007),
		[]byte("this is 32-byte key for Poly1305"),
		[]byte{0xda, 0x84, 0xbc, 0xab, 0x02, 0x67, 0x6c, 0x38, 0xcd, 0xb0, 0x15, 0x60, 0x42, 0x74, 0xc2, 0xaa},
	},
	// empty key results in empty tag irrespective of message contents
	{
		make([]byte, 0),
		make([]byte, 32),
		make([]byte, 16),
	},
	{
		make([]byte, 1),
		make([]byte, 32),
		make([]byte, 16),
	},
	{
		[]byte("Hello world!"),
		make([]byte, 32),
		make([]byte, 16),
	},
	// zero length message
	{
		[]byte{},
		append(mult, sNul...), // as long as S-part is zero, empty result
		make([]byte, 16),
	},
	{
		[]byte{},
		append(mult, sSet...), // when S-part set, get XOR-ed result
		sSet,
	},
	// single zero byte
	{
		[]byte{0x00},
		append(mult, sNul...),
		TagZeroByteMessage_1,
	},
	{
		[]byte{0x00},
		append(mult, sSet...),
		TagZeroByteMessage_2,
	},
	// single byte with first bit set
	{
		[]byte{0x01},
		append(mult, sNul...),
		TagFirstBitMessage_1,
	},
	{
		[]byte{0x01},
		append(mult, sSet...),
		TagFirstBitMessage_2,
	},
	// single byte with second bit set
	{
		[]byte{0x02},
		append(mult, sNul...),
		TagSecndBitMessage_1,
	},
	{
		[]byte{0x02},
		append(mult, sSet...),
		TagSecndBitMessage_2,
	},
	{
		// This test triggers an edge-case. See https://go-review.googlesource.com/#/c/30101/.
		[]byte{0x81, 0xd8, 0xb2, 0xe4, 0x6a, 0x25, 0x21, 0x3b, 0x58, 0xfe, 0xe4, 0x21, 0x3a, 0x2a, 0x28, 0xe9, 0x21, 0xc1, 0x2a, 0x96, 0x32, 0x51, 0x6d, 0x3b, 0x73, 0x27, 0x27, 0x27, 0xbe, 0xcf, 0x21, 0x29},
		[]byte{0x3b, 0x3a, 0x29, 0xe9, 0x3b, 0x21, 0x3a, 0x5c, 0x5c, 0x3b, 0x3b, 0x05, 0x3a, 0x3a, 0x8c, 0x0d},
		[]byte{0x6d, 0xc1, 0x8b, 0x8c, 0x34, 0x4c, 0xd7, 0x99, 0x27, 0x11, 0x8b, 0xbe, 0x84, 0xb7, 0xf3, 0x14},
	},
	// From: https://tools.ietf.org/html/rfc7539#section-2.5.2
	{
		fromHex("43727970746f6772617068696320466f72756d2052657365617263682047726f7570"),
		fromHex("85d6be7857556d337f4452fe42d506a80103808afb0db2fd4abff6af4149f51b"),
		fromHex("a8061dc1305136c6c22b8baf0c0127a9"),
	},
}

func TestVectors(t *testing.T) {
	for i, v := range vectors {
		var key [32]byte

		msg := v.msg
		copy(key[:], v.key)

		out := Sum(msg, key)
		if !bytes.Equal(out[:], v.tag) {
			t.Errorf("Test vector %d : got: %x expected: %x", i, out[:], v.tag)
		}

		h := New(key)
		h.Write(msg)
		tag := h.Sum(nil)
		if !bytes.Equal(tag[:], v.tag) {
			t.Errorf("Test vector %d : got: %x expected: %x", i, tag[:], v.tag)
		}

		var mac [16]byte
		copy(mac[:], v.tag)
		if !Verify(&mac, msg, key) {
			t.Errorf("Test vector %d : Verify failed", i)
		}
	}
}

func TestWriteAfterSum(t *testing.T) {
	msg := make([]byte, 64)
	for i := range msg {
		h := New([32]byte{})

		if _, err := h.Write(msg[:i]); err != nil {
			t.Fatalf("Iteration %d: poly1305.Hash returned unexpected error: %s", i, err)
		}
		h.Sum(nil)
		if _, err := h.Write(nil); err == nil {
			t.Fatalf("Iteration %d: poly1305.Hash returned no error for write after sum", i)
		}
	}
}

func testWrite(t *testing.T, size int) {
	var key [32]byte
	for i := range key {
		key[i] = byte(i)
	}

	h := New(key)

	var msg1 []byte
	msg0 := make([]byte, size)
	for i := range msg0 {
		h.Write(msg0[:i])
		msg1 = append(msg1, msg0[:i]...)
	}

	tag0 := h.Sum(nil)
	tag1 := Sum(msg1, key)

	if !bytes.Equal(tag0[:], tag1[:]) {
		t.Fatalf("Sum differ from poly1305.Sum\n Sum: %s \n poly1305.Sum: %s", hex.EncodeToString(tag0[:]), hex.EncodeToString(tag1[:]))
	}
}

func TestWrite(t *testing.T) {

	for size := 0; size < 128; size++ {
		testWrite(t, size)
	}
}
// Benchmarks

func BenchmarkSum_64(b *testing.B)    { benchmarkSum(b, 64) }
func BenchmarkSum_256(b *testing.B)   { benchmarkSum(b, 256) }
func BenchmarkSum_1K(b *testing.B)    { benchmarkSum(b, 1024) }
func BenchmarkSum_8K(b *testing.B)    { benchmarkSum(b, 8*1024) }
func BenchmarkWrite_64(b *testing.B)  { benchmarkWrite(b, 64) }
func BenchmarkWrite_256(b *testing.B) { benchmarkWrite(b, 256) }
func BenchmarkWrite_1K(b *testing.B)  { benchmarkWrite(b, 1024) }
func BenchmarkWrite_8K(b *testing.B)  { benchmarkWrite(b, 8*1024) }

func benchmarkSum(b *testing.B, size int) {
	var key [32]byte
	for i := range key {
		key[i] = byte(i << 3)
	}

	msg := make([]byte, size)
	for i := range msg {
		msg[i] = byte(i)
	}

	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(msg, key)
	}
}

func benchmarkWrite(b *testing.B, size int) {
	var key [32]byte
	for i := range key {
		key[i] = byte(i << 3)
	}
	h := New(key)

	msg := make([]byte, size)
	for i := range msg {
		msg[i] = byte(i)
	}

	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Write(msg)
	}
}
