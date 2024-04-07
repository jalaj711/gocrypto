package aes

import (
	"reflect"
	"testing"
)

func Test_expandKey(t *testing.T) {
	type args struct {
		key []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "128-bit",
			args: args{
				key: []uint32{0xf1571c9, 0x47d9e859, 0xcb7add6, 0xaf7f6798},
			},
			want: []uint32{
				0xf1571c9, 0x47d9e859, 0xcb7add6, 0xaf7f6798,
				0xdc9037b0, 0x9b49dfe9, 0x97fe723f, 0x388115a7,
				0xd2c96bb7, 0x4980b45e, 0xde7ec661, 0xe6ffd3c6,
				0xc0afdf39, 0x892f6b67, 0x5751ad06, 0xb1ae7ec0,
				0x2c5c65f1, 0xa5730e96, 0xf222a390, 0x438cdd50,
				0x589d36eb, 0xfdee387d, 0xfcc9bed, 0x4c4046bd,
				0x71c74cc2, 0x8c2974bf, 0x83e5ef52, 0xcfa5a9ef,
				0x37149348, 0xbb3de7f7, 0x38d808a5, 0xf77da14a,
				0x48264520, 0xf31ba2d7, 0xcbc3aa72, 0x3cbe0b38,
				0xfd0d42cb, 0xe16e01c, 0xc5d54a6e, 0xf96b4156,
				0xb48ef352, 0xba98134e, 0x7f4d5920, 0x86261876,
			},
		},
		{
			name: "192-bit",
			args: args{
				key: []uint32{0xf1571c9, 0x47d9e859, 0xcb7add6, 0xaf7f6798, 0xf1571c9, 0x47d9e859},
			},
			want: []uint32{
				0x0f1571c9, 0x47d9e859, 0x0cb7add6, 0xaf7f6798, 0x0f1571c9, 0x47d9e859,
				0x3b8eba69, 0x7c575230, 0x70e0ffe6, 0xdf9f987e, 0xd08ae9b7, 0x975301ee,
				0xd4f292e1, 0xa8a5c0d1, 0xd8453f37, 0x07daa749, 0xd7504efe, 0x40034f10,
				0xab7658e8, 0x03d39839, 0xdb96a70e, 0xdc4c0047, 0x0b1c4eb9, 0x4b1f01a9,
				0x630a8b5b, 0x60d91362, 0xbb4fb46c, 0x6703b42b, 0x6c1ffa92, 0x2700fb3b,
				0x10056997, 0x70dc7af5, 0xcb93ce99, 0xac907ab2, 0xc08f8020, 0xe78f7b1b,
				0x4324c603, 0x33f8bcf6, 0xf86b726f, 0x54fb08dd, 0x947488fd, 0x73fbf3e6,
				0x0c29488c, 0x3fd1f47a, 0xc7ba8615, 0x93418ec8, 0x07350635, 0x74cef5d3,
				0x07cf2e1e, 0x381eda64, 0xffa45c71, 0x6ce5d2b9,
			},
		},
		{
			name: "256-bit",
			args: args{
				key: []uint32{0xf1571c9, 0x47d9e859, 0xcb7add6, 0xaf7f6798, 0xf1571c9, 0x47d9e859, 0xcb7add6, 0xaf7f6798},
			},
			want: []uint32{
				0x0f1571c9, 0x47d9e859, 0x0cb7add6, 0xaf7f6798, 0x0f1571c9, 0x47d9e859, 0x0cb7add6, 0xaf7f6798,
				0xdc9037b0, 0x9b49dfe9, 0x97fe723f, 0x388115a7, 0x08192895, 0x4fc0c0cc, 0x43776d1a, 0xec080a82,
				0xeef7247e, 0x75befb97, 0xe24089a8, 0xdac19c0f, 0x5f61f6e3, 0x10a1362f, 0x53d65b35, 0xbfde51b7,
				0xf7268d76, 0x829876e1, 0x60d8ff49, 0xba196346, 0xabb50db9, 0xbb143b96, 0xe8c260a3, 0x571c3114,
				0x63e1772d, 0xe17901cc, 0x81a1fe85, 0x3bb89dc3, 0x49d95397, 0xf2cd6801, 0x1a0f08a2, 0x4d1339b6,
				0x0ef339ce, 0xef8a3802, 0x6e2bc687, 0x55935b44, 0xb5056a8c, 0x47c8028d, 0x5dc70a2f, 0x10d43399,
				0x6630d704, 0x89baef06, 0xe7912981, 0xb20272c5, 0x82722a2a, 0xc5ba28a7, 0x987d2288, 0x88a91111,
				0xf5b255c0, 0x7c08bac6, 0x9b999347, 0x299be182,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expandKey(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
