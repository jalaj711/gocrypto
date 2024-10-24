package sha512

import (
	"reflect"
	"testing"
)

func Test_getWords(t *testing.T) {
	type args struct {
		block []byte
	}
	tests := []struct {
		name string
		args args
		want [80]uint64
	}{
		{
			name: "get-words-sha512-test1",
			args: args{
				block: Pad([]byte("abc")),
			},
			want: [80]uint64{
				0x6162638000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000000,
				0x0000000000000000, 0x0000000000000018,
				0x6162638000000000, 0x00030000000000c0,
				0x0a9699a24c700003, 0x00000c0060000603,
				0x549ef62639858996, 0x00c0003300003c00,
				0x1497007a8a0e9dbc, 0x62e56500cc0780f0,
				0x7760dd475a538797, 0xf1554b711c1c0003,
				0xca2993a4345d9ff2, 0x5e0e66b5c783dd32,
				0xe25a625d00494b62, 0x9f44486fb1e4fbd2,
				0xb31b8c2b06085f2f, 0x0e987660934142f6,
				0xa4af2cfd09fbb924, 0xad289e2e0bd53186,
				0x3c74563aa2f9673e, 0x6ccdcd14cc14b53f,
				0xc3f925b337f22bde, 0x5bcc77a75ad95b54,
				0x3ec2257adca09a52, 0x28246960001fc5eb,
				0x04e33a75ce2be88a, 0x7d5314b3c359e0e7,
				0xaef7a285ff251266, 0x0b8472581deea04f,
				0xb174e26eddc7b033, 0x5d63bae58ddd88de,
				0x4c044007b744ccbb, 0xe6a9aa4d74dc7d43,
				0xebeaf1237248019c, 0x361e80b2d00f3193,
				0x2e9839125df3b175, 0x3319629293ad5363,
				0x9cbc5d89ac1b89d5, 0x275e23ffeeca50b7,
				0x3b80d680bf69ef58, 0x0d0696933945a125,
				0x7533eabcb786ff00, 0xb89826cee6fbf0e5,
				0x249b4fbcad623e9f, 0x4aea9df2b02d6f1e,
				0x2cc57475a55e8d8f, 0xb2574ae938d8be89,
				0xc1b35a57b16d6aea, 0xcc4918b5949206bb,
				0x5099c3add79f90ec, 0x5ea81d78e7660bf1,
				0xebee6267405ac2a9, 0xb01f21926108a4ab,
				0x786433dd2fe65556, 0xc54a6eaa24a0552c,
				0xb3c8f1530bdbaa9e, 0xbb8abfe56f469338,
				0xf63d4265cc1c5a78, 0xbe8355ea73129afb,
				0x49e2db8ebdcfbeb5, 0x82269d4a883a3d99,
				0xfdf53df3011f362b, 0x464af5671d71c12e,
				0xe449b68198ec611c, 0x92aeeed1a7bcf7d2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWords(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngetWords() = %x\nwant       = %x", got, tt.want)
			}
		})
	}
}
