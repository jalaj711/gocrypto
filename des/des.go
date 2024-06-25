package des

type DES struct {
	__roundkeys [16][6]byte
}

func (des *DES) GetBlockSize() int {
	return 8
}

func (des *DES) Init(key []byte) {
	des.__roundkeys = getRoundKeys([8]byte(key[:8]))
}

// Encrypt64 takes 64-bit data block and encrypts it using a 64-bit key by applying DES
func (des *DES) Encrypt64(data [8]byte) (encrypted [8]byte) {
	// IP
	data = initial_permutation(data)

	var L0, R0, L1, R1, out [4]byte
	L0 = [4]byte{data[0], data[1], data[2], data[3]}
	R0 = [4]byte{data[4], data[5], data[6], data[7]}
	for i := 0; i < 16; i++ {
		// Basically through every round we do
		// L(i) = R(i-1)
		// R(i) = L(i-1) XOR F(R(i-1), ROUNDKEY)
		L1 = R0
		out = round_function(R0, des.__roundkeys[i])
		for j := 0; j < 4; j++ {
			R1[j] = L0[j] ^ out[j]
		}
		L0 = L1
		R0 = R1
	}

	// 32-bit swap
	encrypted = [8]byte{R0[0], R0[1], R0[2], R0[3], L0[0], L0[1], L0[2], L0[3]}

	// IP-1
	encrypted = inverse_initial_permutation(encrypted)
	return encrypted
}

// Decrypt64 takes 64-bit data block and decrypts it using a 64-bit key by applying DES
func (des *DES) Decrypt64(data [8]byte) (decrypted [8]byte) {

	data = initial_permutation(data)

	var L0, R0, L1, R1, out [4]byte
	L0 = [4]byte{data[0], data[1], data[2], data[3]}
	R0 = [4]byte{data[4], data[5], data[6], data[7]}

	// decryption is same as encryption except round keys are applied in reverse order
	for i := 15; i >= 0; i-- {
		L1 = R0
		out = round_function(R0, des.__roundkeys[i])
		for j := 0; j < 4; j++ {
			R1[j] = L0[j] ^ out[j]
		}
		L0 = L1
		R0 = R1
	}

	// 32-bit swap
	decrypted = [8]byte{R0[0], R0[1], R0[2], R0[3], L0[0], L0[1], L0[2], L0[3]}

	// IP-1
	decrypted = inverse_initial_permutation(decrypted)
	return decrypted
}

// Encrypt encrypts a given byte array of 8 bytes
func (des *DES) Encrypt(data []byte) (encrypted []byte) {
	if len(data) != 8 {
		panic("plaintext must be of 64-bits")
	}
	encrypted = make([]byte, 0, len(data))
	block := des.Encrypt64(([8]byte)(data))
	encrypted = append(encrypted, block[:]...)
	return encrypted
}

// Decrypt decrypts a given byte array that was encrypted using DES 64-bit encryption
func (des *DES) Decrypt(data []byte) (decrypted []byte) {
	if len(data) != 8 {
		panic("ciphertext must be of 64-bits")
	}
	decrypted = make([]byte, 0)
	temp := des.Decrypt64(([8]byte)(data))
	decrypted = append(decrypted, temp[:]...)
	return decrypted
}
