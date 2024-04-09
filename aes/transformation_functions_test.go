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

func Test_addRoundKey(t *testing.T) {
	type args struct {
		state    []uint32
		roundKey []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "T1",
			args: args{
				state:    []uint32{0x4740a34c, 0x37d4709f, 0x94e43a42, 0xeda5a6bc},
				roundKey: []uint32{0xac192857, 0x77fad15c, 0x66dc2900, 0xf321416a},
			},
			want: []uint32{0xeb598b1b, 0x402ea1c3, 0xf2381342, 0x1e84e7d6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addRoundKey(tt.args.state, tt.args.roundKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addRoundKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test__multiply(t *testing.T) {
	type args struct {
		num        byte
		multiplyBy byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "0x87 x 0x02",
			args: args{
				num:        0xae,
				multiplyBy: 0x02,
			},
			want: 0x47,
		},
		{
			name: "0x57 x 0x02",
			args: args{
				num:        0x57,
				multiplyBy: 0x02,
			},
			want: 0xae,
		},
		{
			name: "0x6e x 0x03",
			args: args{
				num:        0x6e,
				multiplyBy: 0x03,
			},
			want: 0xb2,
		},
		{
			name: "0x25 x 0x03",
			args: args{
				num:        0x25,
				multiplyBy: 0x03,
			},
			want: 0x6f,
		},
		{
			name: "0x57 x 0x04",
			args: args{
				num:        0x57,
				multiplyBy: 0x04,
			},
			want: 0x47,
		},
		{
			name: "0x57 x 0x08",
			args: args{
				num:        0x57,
				multiplyBy: 0x08,
			},
			want: 0x8e,
		},
		{
			name: "0x57 x 0x10",
			args: args{
				num:        0x57,
				multiplyBy: 0x10,
			},
			want: 0x07,
		},
		{
			name: "0x57 x 0x20",
			args: args{
				num:        0x57,
				multiplyBy: 0x20,
			},
			want: 0x0e,
		},
		{
			name: "0x57 x 0x40",
			args: args{
				num:        0x57,
				multiplyBy: 0x40,
			},
			want: 0x1c,
		},
		{
			name: "0x57 x 0x09",
			args: args{
				num:        0x57,
				multiplyBy: 0x09,
			},
			want: 0xd9,
		},
		{
			name: "0x57 x 0x13",
			args: args{
				num:        0x57,
				multiplyBy: 0x13,
			},
			want: 0xfe,
		},
		{
			name: "0xee x 0xd",
			args: args{
				num:        0xee,
				multiplyBy: 0xd,
			},
			want: 0x4a,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _multiply(tt.args.num, tt.args.multiplyBy); got != tt.want {
				t.Errorf("_multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mixColumns(t *testing.T) {
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
				state: []uint32{0x01c6d497, 0x01c6d4ec, 0x01c6d4c3, 0x01c6d595},
			},
			want: []uint32{0x01c6d54c, 0x01c6d59f, 0x01c6d742, 0x01c6d6bc},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mixColumns(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mixColumns() = %x, want %x", got, tt.want)
			}
		})
	}
}
