package aes

import (
	"reflect"
	"testing"
)

func Test_subBytes(t *testing.T) {
	type args struct {
		state []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "T1",
			args: args{
				state: []uint32{0xea046585, 0x83455d96, 0x5c3398b0, 0xf02dadc5},
			},
			want: []uint32{0x87f24d97, 0xec6e4c90, 0x4ac346e7, 0x8cd895a6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subBytes(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shiftRows(t *testing.T) {
	type args struct {
		state []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "T1",
			args: args{
				state: []uint32{0xea046585, 0x83455d96, 0x5c3398b0, 0xf02dadc5},
			},
			want: []uint32{0xea046585, 0x455d9683, 0x98b05c33, 0xc5f02dad},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftRows(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
