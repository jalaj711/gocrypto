package des

import "errors"

// Encrypt64 takes 64-bit data block and encrypts it using a 64-bit key by applying DES
func Encrypt64(data [8]byte, key [8]byte) (encrypted [8]byte) {
	round_keys := getRoundKeys(key)

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
		out = round_function(R0, round_keys[i])
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
func Decrypt64(data [8]byte, key [8]byte) (decrypted [8]byte) {
	round_keys := getRoundKeys(key)

	data = initial_permutation(data)

	var L0, R0, L1, R1, out [4]byte
	L0 = [4]byte{data[0], data[1], data[2], data[3]}
	R0 = [4]byte{data[4], data[5], data[6], data[7]}

	// decryption is same as encryption except round keys are applied in reverse order
	for i := 15; i >= 0; i-- {
		L1 = R0
		out = round_function(R0, round_keys[i])
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

// Encrypt encrypts a given byte array of any number of bytes using the 64-bit key supplied
// the input message is first padded to make it a multiple of 64-bits
// the encryption happens using electronic codebook mode
func Encrypt(data []byte, key [8]byte) (encrypted []byte) {
	data = addPadding(data)
	encrypted = make([]byte, 0, len(data))
	var block [8]byte
	for i := 0; i < len(data); i += 8 {
		block = Encrypt64(([8]byte)(data[i:i+8]), key)
		encrypted = append(encrypted, block[:]...)
	}
	return encrypted
}

// Decrypt decrypts a given byte array that was encrypted using DES 64-bit electronic codebook mode encryption using CMS padding
// the padding is removed before returning the plaintext
func Decrypt(data []byte, key [8]byte) (decrypted []byte, err error) {
	if len(data)%8 != 0 {
		return decrypted, errors.New("ciphertext must be a multiple of 64-bits")
	}
	decrypted = make([]byte, 0, len(data))
	var block [8]byte
	for i := 0; i < len(data); i += 8 {
		block = Decrypt64(([8]byte)(data[i:i+8]), key)
		decrypted = append(decrypted, block[:]...)
	}
	return removePadding(decrypted)
}
