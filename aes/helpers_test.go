package aes

import "testing"

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
