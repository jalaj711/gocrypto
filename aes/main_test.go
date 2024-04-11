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
			name: "T1",
			args: args{
				data: []uint32{0x01234567, 0x89abcdef, 0xfedcba98, 0x76543210}, key: []uint32{0x0f1571c9, 0x47d9e859, 0x0cb7add6, 0xaf7f6798},
			},
			want: []uint32{0xff0b844a, 0x0853bf7c, 0x6934ab43, 0x64148fb9},
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
