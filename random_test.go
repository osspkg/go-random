/*
 *  Copyright (c) 2024-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package random

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnit_Bytes(t *testing.T) {
	size := 10
	r1 := Bytes(size)
	r2 := Bytes(size)

	fmt.Println(string(r1), string(r2))

	if len(r1) != size || len(r2) != size {
		t.Errorf("invalid len, is not %d", size)
	}
	if reflect.DeepEqual(r1, r2) {
		t.Errorf("result is not random")
	}
}

func TestUnit_String(t *testing.T) {
	size := 10
	r1 := String(size)
	r2 := String(size)

	fmt.Println(string(r1), string(r2))

	if len(r1) != size || len(r2) != size {
		t.Errorf("invalid len, is not %d", size)
	}
	if reflect.DeepEqual(r1, r2) {
		t.Errorf("result is not random")
	}
}

func TestUnit_BytesOf(t *testing.T) {
	size := 10
	src := []byte("1234567890")
	r1 := BytesOf(size, src)
	r2 := BytesOf(size, src)

	fmt.Println(string(r1), string(r2))

	if len(r1) != size || len(r2) != size {
		t.Errorf("invalid len, is not %d", size)
	}
	if reflect.DeepEqual(r1, r2) {
		t.Errorf("result is not random")
	}
}

/*
goos: linux
goarch: amd64
pkg: go.osspkg.com/random
cpu: 12th Gen Intel(R) Core(TM) i9-12900KF
Benchmark_Bytes64
Benchmark_Bytes64-24         	10615876	       114.5 ns/op	     216 B/op	       3 allocs/op
Benchmark_Bytes256
Benchmark_Bytes256-24        	 5578935	       208.9 ns/op	     409 B/op	       3 allocs/op
Benchmark_CryptoBytes
Benchmark_CryptoBytes-24     	27048033	        43.04 ns/op	      64 B/op	       1 allocs/op
Benchmark_CryptoBase64
Benchmark_CryptoBase64-24    	16419511	        74.44 ns/op	     256 B/op	       3 allocs/op
*/

func Benchmark_Bytes64(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Bytes(64)
		}
	})
}

func Benchmark_Bytes256(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Bytes(256)
		}
	})
}

func Benchmark_CryptoBytes(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CryptoBytes(64)
		}
	})
}

func Benchmark_CryptoBase64(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CryptoBase64(64)
		}
	})
}
