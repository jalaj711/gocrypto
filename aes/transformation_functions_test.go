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

func Test_invSubBytes(t *testing.T) {
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
				state: []uint32{0x87f24d97, 0xec6e4c90, 0x4ac346e7, 0x8cd895a6},
			},
			want: []uint32{0xea046585, 0x83455d96, 0x5c3398b0, 0xf02dadc5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invSubBytes(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invSubBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invShiftRows(t *testing.T) {
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
				state: []uint32{0xea046585, 0x455d9683, 0x98b05c33, 0xc5f02dad},
			},
			want: []uint32{0xea046585, 0x83455d96, 0x5c3398b0, 0xf02dadc5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invShiftRows(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invShiftRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_invMixColumns(t *testing.T) {
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
				state: []uint32{0x01c6d54c, 0x01c6d59f, 0x01c6d742, 0x01c6d6bc},
			},
			want: []uint32{0x01c6d497, 0x01c6d4ec, 0x01c6d4c3, 0x01c6d595},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invMixColumns(tt.args.state); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invMixColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}
