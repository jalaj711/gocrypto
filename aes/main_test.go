package aes

import (
	"reflect"
	"testing"
)

func TestEncrypt128(t *testing.T) {
	type args struct {
		data []uint32
		key  []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "T1-128",
			args: args{
				data: []uint32{0x01234567, 0x89abcdef, 0xfedcba98, 0x76543210}, key: []uint32{0x0f1571c9, 0x47d9e859, 0x0cb7add6, 0xaf7f6798},
			},
			want: []uint32{0xff0b844a, 0x0853bf7c, 0x6934ab43, 0x64148fb9},
		},
		{
			name: "T2-192",
			args: args{
				data: []uint32{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff}, key: []uint32{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f, 0x10111213, 0x14151617},
			},
			want: []uint32{0xdda97ca4, 0x864cdfe0, 0x6eaf70a0, 0xec0d7191},
		},
		{
			name: "T3-256",
			args: args{
				data: []uint32{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff}, key: []uint32{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f, 0x10111213, 0x14151617, 0x18191a1b, 0x1c1d1e1f},
			},
			want: []uint32{0x8ea2b7ca, 0x516745bf, 0xeafc4990, 0x4b496089},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aes := new(AES)
			aes.Init(_uintArrToByte(tt.args.key))
			if got := aes.Encrypt128(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt128() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestDecrypt128(t *testing.T) {
	type args struct {
		data []uint32
		key  []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "T1-128",
			args: args{
				data: []uint32{0xff0b844a, 0x0853bf7c, 0x6934ab43, 0x64148fb9}, key: []uint32{0x0f1571c9, 0x47d9e859, 0x0cb7add6, 0xaf7f6798},
			},
			want: []uint32{0x01234567, 0x89abcdef, 0xfedcba98, 0x76543210},
		},
		{
			name: "T2-192",
			args: args{
				data: []uint32{0xdda97ca4, 0x864cdfe0, 0x6eaf70a0, 0xec0d7191}, key: []uint32{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f, 0x10111213, 0x14151617},
			},
			want: []uint32{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff},
		},
		{
			name: "T3-256",
			args: args{
				data: []uint32{0x8ea2b7ca, 0x516745bf, 0xeafc4990, 0x4b496089}, key: []uint32{0x00010203, 0x04050607, 0x08090a0b, 0x0c0d0e0f, 0x10111213, 0x14151617, 0x18191a1b, 0x1c1d1e1f},
			},
			want: []uint32{0x00112233, 0x44556677, 0x8899aabb, 0xccddeeff},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aes := new(AES)
			aes.Init(_uintArrToByte(tt.args.key))
			if got := aes.Decrypt128(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt128() = %x, want %x", got, tt.want)
			}
		})
	}
}
