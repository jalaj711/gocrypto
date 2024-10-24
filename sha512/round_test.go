package sha512

import (
	"reflect"
	"testing"
)

func Test_roundF(t *testing.T) {
	type args struct {
		W uint64
		H *registers
		t int
	}
	tests := []struct {
		name string
		args args
		want registers
	}{
		{
			name: "test-round-func-1",
			args: args{
				W: 0x6162638000000000,
				H: &registers{
					a: 0x6A09E667F3BCC908,
					b: 0xBB67AE8584CAA73B,
					c: 0x3C6EF372FE94F82B,
					d: 0xA54FF53A5F1D36F1,
					e: 0x510E527FADE682D1,
					f: 0x9B05688C2B3E6C1F,
					g: 0x1F83D9ABFB41BD6B,
					h: 0x5BE0CD19137E2179,
				},
				t: 0,
			},
			want: registers{
				a: 0xF6AFCEB8BCFCDDF5,
				b: 0x6A09E667F3BCC908,
				c: 0xBB67AE8584CAA73B,
				d: 0x3C6EF372FE94F82B,
				e: 0x58CB02347AB51F91,
				f: 0x510E527FADE682D1,
				g: 0x9B05688C2B3E6C1F,
				h: 0x1F83D9ABFB41BD6B,
			},
		},
		{
			name: "test-round-func-2",
			args: args{
				W: 0x0,
				H: &registers{
					a: 0xF6AFCEB8BCFCDDF5,
					b: 0x6A09E667F3BCC908,
					c: 0xBB67AE8584CAA73B,
					d: 0x3C6EF372FE94F82B,
					e: 0x58CB02347AB51F91,
					f: 0x510E527FADE682D1,
					g: 0x9B05688C2B3E6C1F,
					h: 0x1F83D9ABFB41BD6B,
				},
				t: 1,
			},

			want: registers{
				a: 0x1320F8C9FB872CC0,
				b: 0xF6AFCEB8BCFCDDF5,
				c: 0x6A09E667F3BCC908,
				d: 0xBB67AE8584CAA73B,
				e: 0xC3D4EBFD48650FFA,
				f: 0x58CB02347AB51F91,
				g: 0x510E527FADE682D1,
				h: 0x9B05688C2B3E6C1F,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roundF(tt.args.W, tt.args.H, tt.args.t)
			if !reflect.DeepEqual(*tt.args.H, tt.want) {
				t.Errorf("roundF() = %v, want %v", tt.args.H, tt.want)
			}
		})
	}
}
