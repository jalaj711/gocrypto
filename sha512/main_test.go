package sha512

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "hash-test-1-single-block",
			args: args{
				input: []byte("abc"),
			},
			want: []uint64{
				0xDDAF35A193617ABA,
				0xCC417349AE204131,
				0x12E6FA4E89A97EA2,
				0x0A9EEEE64B55D39A,
				0x2192992A274FC1A8,
				0x36BA3C23A3FEEBBD,
				0x454D4423643CE80E,
				0x2A9AC94FA54CA49F,
			},
		},
		{
			name: "hash-test-2-double-block",
			args: args{
				input: []byte("abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu"),
			},
			want: []uint64{
				0x8E959B75DAE313DA,
				0x8CF4F72814FC143F,
				0x8F7779C6EB9F7FA1,
				0x7299AEADB6889018,
				0x501D289E4900F7E4,
				0x331B99DEC4B5433A,
				0xC7D329EEB6DD2654,
				0x5E96E55B874BE909,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hash() = %x, want %x", got, tt.want)
			}
		})
	}
}
