package modes

type CFB struct {
	cipher  BlockCipher
	padding Padding
	iv      []byte
}

func (cfb *CFB) Init(key []byte, iv []byte) {
	cfb.cipher.Init(key)
	if len(iv) > 0 {
		cfb.iv = iv
	}
	if len(cfb.iv) != cfb.cipher.GetBlockSize() {
		panic("Invalid IV length. Must be equal to block size.")
	}
}

func (cfb *CFB) Encrypt(data []byte) []byte {
	blockSize := cfb.cipher.GetBlockSize()
	plaintext := cfb.padding.Pad(data, blockSize)
	encrypted := make([]byte, 0, len(plaintext))
	iv := make([]byte, blockSize)
	copy(iv, cfb.iv)
	var temp byte
	for i := 0; i < len(plaintext); i += 1 {
		iv = cfb.cipher.Encrypt(iv)
		temp = plaintext[i] ^ iv[0]
		iv = append(iv[1:], temp)
		encrypted = append(encrypted, temp)
	}
	return encrypted
}

func (cfb *CFB) Decrypt(data []byte) []byte {
	blockSize := cfb.cipher.GetBlockSize()
	if len(data)%blockSize != 0 {
		panic("Invalid ciphertext")
	}
	plaintext := make([]byte, 0, len(data))
	iv := make([]byte, blockSize)
	copy(iv, cfb.iv)
	for i := 0; i < len(data); i += 1 {
		iv = cfb.cipher.Encrypt(iv)
		plaintext = append(plaintext, data[i]^iv[0])
		iv = append(iv[1:], data[i])
	}
	return cfb.padding.UnPad(plaintext, blockSize)
}
