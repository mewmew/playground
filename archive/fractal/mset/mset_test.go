package mset

import (
	"bytes"
	"image/png"
	"testing"
)

func BenchmarkAt(b *testing.B) {
	buf := new(bytes.Buffer)
	img := New(128, 128)
	for i := 0; i < b.N; i++ {
		err := png.Encode(buf, img)
		if err != nil {
			b.Error(err)
		}
		buf.Reset()
	}
}
