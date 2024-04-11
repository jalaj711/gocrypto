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

func mixColumns(state []uint32) []uint32 {
	s := [4][4]byte{}
	ss := [4][4]byte{}

	for i := 0; i < 4; i++ {
		s[i][0] = byte(state[i] >> 24)
		s[i][1] = byte(state[i] >> 16)
		s[i][2] = byte(state[i] >> 8)
		s[i][3] = byte(state[i])
	}
	for i := 0; i < 4; i++ {
		ss[0][i] = _multiply(s[0][i], 2) ^ _multiply(s[1][i], 3) ^ s[2][i] ^ s[3][i]
		ss[1][i] = _multiply(s[1][i], 2) ^ _multiply(s[2][i], 3) ^ s[3][i] ^ s[0][i]
		ss[2][i] = _multiply(s[2][i], 2) ^ _multiply(s[3][i], 3) ^ s[0][i] ^ s[1][i]
		ss[3][i] = _multiply(s[3][i], 2) ^ _multiply(s[0][i], 3) ^ s[1][i] ^ s[2][i]
	}
	for i := 0; i < 4; i++ {
		state[i] = uint32(ss[i][0])<<24 | uint32(ss[i][1])<<16 | uint32(ss[i][2])<<8 | uint32(ss[i][3])
	}
	return state
}
