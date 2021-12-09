// Code generated by ch-gen-int, DO NOT EDIT.

package proto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColUInt64_DecodeColumn(t *testing.T) {
	const rows = 50_000
	var data ColUInt64
	for i := 0; i < rows; i++ {
		data = append(data, uint64(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	br := bytes.NewReader(buf.Buf)
	r := NewReader(br)

	var dec ColUInt64
	require.NoError(t, dec.DecodeColumn(r, rows))
	require.Equal(t, data, dec)
}

func BenchmarkColUInt64_DecodeColumn(b *testing.B) {
	const rows = 50_000
	var data ColUInt64
	for i := 0; i < rows; i++ {
		data = append(data, uint64(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	br := bytes.NewReader(buf.Buf)
	r := NewReader(br)

	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	var dec ColUInt64
	for i := 0; i < b.N; i++ {
		br.Reset(buf.Buf)
		r.s.Reset(br)
		dec.Reset()

		if err := dec.DecodeColumn(r, rows); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColUInt64_EncodeColumn(b *testing.B) {
	const rows = 50_000
	var data ColUInt64
	for i := 0; i < rows; i++ {
		data = append(data, uint64(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		data.EncodeColumn(&buf)
	}
}
