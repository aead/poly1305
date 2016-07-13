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

var vectors = []struct {
	key, msg, tag string
}{
	// From: https://tools.ietf.org/html/rfc7539#section-2.5.2
	{
		key: "85d6be7857556d337f4452fe42d506a80103808afb0db2fd4abff6af4149f51b",
		msg: "43727970746f6772617068696320466f72756d2052657365617263682047726f7570",
		tag: "a8061dc1305136c6c22b8baf0c0127a9",
	},
}

func TestVectors(t *testing.T) {
	for i, v := range vectors {
		key := fromHex(v.key)
		msg := fromHex(v.msg)
		tag := fromHex(v.tag)

		var sum [TagSize]byte
		var k [32]byte

		copy(k[:], key)
		Sum(&sum, msg, &k)
		if !bytes.Equal(sum[:], tag) {
			t.Fatalf("Test vector %d : Poly1305 Tags are not equal:\nFound:    %v\nExpected: %v", i, sum, tag)
		}

		p := New(&k)
		p.Write(msg)
		p.Sum(&sum)
		if !bytes.Equal(sum[:], tag) {
			t.Fatalf("Test vector %d : Poly1305 Tags are not equal:\nFound:    %v\nExpected: %v", i, sum[:], tag)
		}
	}
}
