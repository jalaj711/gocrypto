package aes

// helper function to perform multiplication in GF(256)
func _multiply(num byte, multiplyBy byte) byte {
	if multiplyBy&1 == 1 {
		return num ^ _multiply(num, multiplyBy^1)
	}
	ans := num << 1
	if num&0x80 == 0x80 {
		ans ^= 0x1b
	}
	multiplyBy = multiplyBy >> 1
	if multiplyBy > 1 {
		return _multiply(ans, multiplyBy)
	}
	return ans
}

func subWord(word uint32) (result uint32) {
	result = uint32(SBOX[word>>28][(word>>24)&15]) << 24
	result |= uint32(SBOX[(word>>20)&15][(word>>16)&15]) << 16
	result |= uint32(SBOX[(word>>12)&15][(word>>8)&15]) << 8
	result |= uint32(SBOX[(word>>4)&15][(word)&15])
	return result
}

func subBytes(state []uint32) []uint32 {
	for i := 0; i < len(state); i++ {
		state[i] = subWord(state[i])
	}
	return state
}

func shiftRows(state []uint32) []uint32 {
	for i := 0; i < len(state); i++ {
		state[i] = (state[i] << (i * 8)) | (state[i] >> (32 - i*8))
	}
	return state
}

func addRoundKey(state, roundKey []uint32) []uint32 {
	for i := 0; i < len(state); i++ {
		state[i] = state[i] ^ roundKey[i]
	}
	return state
}
