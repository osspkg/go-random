/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package random

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/rand/v2"
	"sync"
	"time"
)

var (
	poolRnd = sync.Pool{New: func() any {
		return createRand()
	}}
	poolDigest = sync.Pool{New: func() any {
		return createDigest(createRand(), 128)
	}}
)

func createRand() *rand.Rand {
	seed2 := uint64(time.Now().UnixNano())
	return rand.New(rand.NewPCG(seed2/100, seed2)) //nolint:gosec
}

func createDigest(rnd *rand.Rand, n int) []byte {
	b := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		v := rnd.IntN(255)
		b = append(b, byte(v))
	}
	return b
}

func Rand(call func(*rand.Rand)) {
	rnd := poolRnd.Get().(*rand.Rand)
	defer poolRnd.Put(rnd)

	call(rnd)
}

func BytesOf(n int, src []byte) []byte {
	rnd := poolRnd.Get().(*rand.Rand)
	defer poolRnd.Put(rnd)

	tmp := make([]byte, len(src))
	copy(tmp, src)

	rnd.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})

	b := make([]byte, n)
	for i := range b {
		b[i] = tmp[rnd.IntN(len(tmp))]
	}
	return b
}

func StringOf(n int, src string) string {
	return string(BytesOf(n, []byte(src)))
}

func Bytes(n int) []byte {
	digest := poolDigest.Get().([]byte)
	poolDigest.Put(digest) //nolint:staticcheck

	return BytesOf(n, digest)
}

func String(n int) string {
	b := Bytes(n)
	s := base64.StdEncoding.EncodeToString(b)
	if len(s) > n {
		return s[:n]
	}
	return s
}

func Shuffle[T any](v []T) []T {
	rnd := poolRnd.Get().(*rand.Rand)
	defer poolRnd.Put(rnd)

	rnd.Shuffle(len(v), func(i, j int) { v[i], v[j] = v[j], v[i] })
	return v
}

func CryptoBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < 10; i++ {
		if _, err := crand.Read(b); err != nil {
			continue
		}
		return b
	}
	return Bytes(n)
}

func CryptoBase64(n int) string {
	b := CryptoBytes(n)
	return base64.StdEncoding.EncodeToString(b)
}
