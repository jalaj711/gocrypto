package modes

type OFB struct {
	cipher  BlockCipher
	padding Padding
	iv      []byte
}

func (ofb *OFB) Init(key []byte, iv []byte) {
	ofb.cipher.Init(key)
	if len(iv) > 0 {
		ofb.iv = iv
	}
	if len(ofb.iv) != ofb.cipher.GetBlockSize() {
		panic("Invalid IV length. Must be equal to block size.")
	}
}

func (ofb *OFB) Encrypt(data []byte) []byte {
	blockSize := ofb.cipher.GetBlockSize()
	plaintext := ofb.padding.Pad(data, blockSize)
	encrypted := make([]byte, 0, len(plaintext))
	iv := make([]byte, blockSize)
	copy(iv, ofb.iv)
	temp := make([]byte, blockSize)
	for i := 0; i < len(plaintext); i += blockSize {
		iv = ofb.cipher.Encrypt(iv)
		for j := 0; j < blockSize; j++ {
			temp[j] = iv[j] ^ plaintext[i+j]
		}
		encrypted = append(encrypted, temp...)
	}
	return encrypted
}

func (ofb *OFB) Decrypt(data []byte) []byte {
	blockSize := ofb.cipher.GetBlockSize()
	if len(data)%blockSize != 0 {
		panic("Invalid ciphertext")
	}
	plaintext := make([]byte, 0, len(data))
	iv := make([]byte, blockSize)
	copy(iv, ofb.iv)
	temp := make([]byte, blockSize)
	for i := 0; i < len(data); i += blockSize {
		iv = ofb.cipher.Encrypt(iv)
		for j := 0; j < blockSize; j++ {
			temp[j] = iv[j] ^ data[i+j]
		}
		plaintext = append(plaintext, temp...)
	}
	return ofb.padding.UnPad(plaintext, blockSize)
}
