package aes

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
