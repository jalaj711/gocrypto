package modes

import "fmt"

type CBC struct {
	cipher  BlockCipher
	padding Padding
	iv      []byte
}

func (cbc *CBC) Init(key []byte, iv []byte) {
	cbc.cipher.Init(key)
	if len(iv) > 0 {
		cbc.iv = iv
	}
	if len(cbc.iv) != cbc.cipher.GetBlockSize() {
		panic("Invalid IV length. Must be equal to block size.")
	}
}

func (cbc *CBC) Encrypt(data []byte) []byte {
	blockSize := cbc.cipher.GetBlockSize()
	plaintext := cbc.padding.Pad(data, blockSize)
	fmt.Printf("%x", plaintext)
	iv := make([]byte, blockSize)
	copy(iv, cbc.iv)
	for i := 0; i < len(plaintext); i += blockSize {
		for j := 0; j < blockSize; j++ {
			iv[j] = iv[j] ^ plaintext[i+j]
		}
		iv = cbc.cipher.Encrypt(iv)
	}
	return iv
}
