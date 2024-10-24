package sha512

import "testing"

func Test_rotr(t *testing.T) {
	type args struct {
		x uint64
		n int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "T1",
			args: args{
				x: 0x1122334455667788,
				n: 4,
			},
			want: 0x8112233445566778,
		},
		{
			name: "T2",
			args: args{
				x: 0x1122334455667788,
				n: 8,
			},
			want: 0x8811223344556677,
		},
		{
			name: "T3",
			args: args{
				x: 0x1122334455667788,
				n: 64,
			},
			want: 0x1122334455667788,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotr(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("rotr() = %x, want %x", got, tt.want)
			}
		})
	}
}
