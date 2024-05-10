package modes

type ECB struct {
	cipher  BlockCipher
	padding Padding
}

func (ecb *ECB) Init(key []byte) {
	ecb.cipher.Init(key)
}

func (ecb *ECB) Encrypt(data []byte) []byte {
	blockSize := ecb.cipher.GetBlockSize()
	plaintext := ecb.padding.Pad(data, blockSize)
	encrypted := make([]byte, 0, len(plaintext))
	for i := 0; i < len(plaintext); i += blockSize {
		encrypted = append(encrypted, ecb.cipher.Encrypt(plaintext[i:i+blockSize])...)
	}
	return encrypted
}

func (ecb *ECB) Decrypt(data []byte) []byte {
	blockSize := ecb.cipher.GetBlockSize()
	if len(data)%blockSize != 0 {
		panic("Invalid ciphertext")
	}
	plaintext := make([]byte, 0, len(data))
	for i := 0; i < len(data); i += blockSize {
		plaintext = append(plaintext, ecb.cipher.Decrypt(data[i:i+blockSize])...)
	}
	return ecb.padding.UnPad(plaintext, blockSize)
}
