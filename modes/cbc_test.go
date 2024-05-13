package modes

import (
	"github.com/jalaj711/gocrypto/aes"
	"reflect"
	"testing"
)

func TestCBC_Encrypt(t *testing.T) {
	type fields struct {
		cipher  BlockCipher
		padding Padding
		iv      []byte
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: "cbc-aes-1",
			fields: fields{
				cipher:  &aes.AES{},
				padding: &PKCS7{},
				iv: []byte{
					0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0,
					0x0, 0x0, 0x0, 0x0,
				},
			},
			args: args{
				data: []byte{
					0x00, 0x7e, 0x65, 0x77,
					0xe6, 0x57, 0x7a, 0xbc,
					0x62, 0xcd, 0x28, 0x73,
					0x65, 0xf1, 0x99,
				},
			},
			want: []byte{
				0xff, 0x46, 0x30, 0xf2,
				0x5b, 0xbe, 0xb6, 0x3c,
				0x7f, 0xcf, 0x45, 0x76,
				0xf5, 0x0d, 0x4e, 0xa9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cbc := &CBC{
				cipher:  tt.fields.cipher,
				padding: tt.fields.padding,
				iv:      tt.fields.iv,
			}
			cbc.Init([]byte{
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			}, []byte{})
			if got := cbc.Encrypt(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() = %x, want %x", got, tt.want)
			}
		})
	}
}
