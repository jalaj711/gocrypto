package des

type TripleDES128 struct {
	k1 DES
	k2 DES
}
type TripleDES192 struct {
	k1 DES
	k2 DES
	k3 DES
}

func (tdes *TripleDES128) Init(key []byte) {
	tdes.k1 = DES{}
	tdes.k1.Init(key[:8])
	tdes.k2 = DES{}
	tdes.k2.Init(key[8:])
}

func (tdes *TripleDES128) GetBlockSize() int {
	return 8
}

func (tdes *TripleDES128) Encrypt64(plaintext [8]byte) (ciphertext [8]byte) {
	return tdes.k1.Encrypt64(tdes.k2.Decrypt64(tdes.k1.Encrypt64(plaintext)))
}

func (tdes *TripleDES128) Decrypt64(plaintext [8]byte) (ciphertext [8]byte) {
	return tdes.k1.Decrypt64(tdes.k2.Encrypt64(tdes.k1.Decrypt64(plaintext)))
}

// Encrypt encrypts a given byte array of 8 bytes
func (tdes *TripleDES128) Encrypt(data []byte) (encrypted []byte) {
	if len(data) != 8 {
		panic("plaintext must be of 64-bits")
	}
	encrypted = make([]byte, 0, len(data))
	block := tdes.Encrypt64(([8]byte)(data))
	encrypted = append(encrypted, block[:]...)
	return encrypted
}

// Decrypt decrypts a given byte array that was encrypted using DES 64-bit encryption
func (tdes *TripleDES128) Decrypt(data []byte) (decrypted []byte) {
	if len(data) != 8 {
		panic("ciphertext must be of 64-bits")
	}
	decrypted = make([]byte, 0)
	temp := tdes.Decrypt64(([8]byte)(data))
	decrypted = append(decrypted, temp[:]...)
	return decrypted
}

func (tdes *TripleDES192) Init(key []byte) {
	tdes.k1 = DES{}
	tdes.k1.Init(key[:8])
	tdes.k2 = DES{}
	tdes.k2.Init(key[8:16])
	tdes.k3 = DES{}
	tdes.k3.Init(key[16:])
}

func (tdes *TripleDES192) GetBlockSize() int {
	return 8
}

func (tdes *TripleDES192) Encrypt64(plaintext [8]byte) (ciphertext [8]byte) {
	return tdes.k3.Encrypt64(tdes.k2.Decrypt64(tdes.k1.Encrypt64(plaintext)))
}

func (tdes *TripleDES192) Decrypt64(plaintext [8]byte) (ciphertext [8]byte) {
	return tdes.k3.Decrypt64(tdes.k2.Encrypt64(tdes.k1.Decrypt64(plaintext)))
}

// Encrypt encrypts a given byte array of 8 bytes
func (tdes *TripleDES192) Encrypt(data []byte) (encrypted []byte) {
	if len(data) != 8 {
		panic("plaintext must be of 64-bits")
	}
	encrypted = make([]byte, 0, len(data))
	block := tdes.Encrypt64(([8]byte)(data))
	encrypted = append(encrypted, block[:]...)
	return encrypted
}

// Decrypt decrypts a given byte array that was encrypted using DES 64-bit encryption
func (tdes *TripleDES192) Decrypt(data []byte) (decrypted []byte) {
	if len(data) != 8 {
		panic("ciphertext must be of 64-bits")
	}
	decrypted = make([]byte, 0)
	temp := tdes.Decrypt64(([8]byte)(data))
	decrypted = append(decrypted, temp[:]...)
	return decrypted
}
