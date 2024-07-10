/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package random

import (
	crand "crypto/rand"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var (
	digest = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-+=~*@#$%&?!<>")
)

func BytesOf(n int, src []byte) []byte {
	tmp := make([]byte, len(src))
	copy(tmp, src)
	rnd.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})
	b := make([]byte, n)
	for i := range b {
		b[i] = tmp[rnd.Intn(len(tmp))]
	}
	return b
}

func StringOf(n int, src string) string {
	return string(BytesOf(n, []byte(src)))
}

func Bytes(n int) []byte {
	return BytesOf(n, digest)
}

func String(n int) string {
	return string(Bytes(n))
}

func Shuffle(v []string) []string {
	rnd.Shuffle(len(v), func(i, j int) { v[i], v[j] = v[j], v[i] })
	return v
}

func CryptoBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	var err error
	for i := 0; i < 10; i++ {
		if _, err = crand.Read(b); err != nil {
			continue
		}
		return b, nil
	}
	return nil, err
}
