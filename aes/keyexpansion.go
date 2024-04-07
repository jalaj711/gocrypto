package aes

var RCON = []uint32{0x01 << 24, 0x02 << 24, 0x04 << 24, 0x08 << 24, 0x10 << 24, 0x20 << 24, 0x40 << 24, 0x80 << 24, 0x1B << 24, 0x36 << 24}

func subWord(word uint32) (result uint32) {
	result = uint32(SBOX[word>>28][(word>>24)&15]) << 24
	result |= uint32(SBOX[(word>>20)&15][(word>>16)&15]) << 16
	result |= uint32(SBOX[(word>>12)&15][(word>>8)&15]) << 8
	result |= uint32(SBOX[(word>>4)&15][(word)&15])
	return result
}

func expandKey(key []uint32) []uint32 {
	var expandedWords []uint32
	N := len(key)
	var R int
	if N == 4 {
		expandedWords = make([]uint32, 44)
		R = 11
	} else if N == 6 {
		expandedWords = make([]uint32, 52)
		R = 13
	} else if N == 8 {
		expandedWords = make([]uint32, 60)
		R = 15
	} else {
		panic("Length of key should only be 4, 6 or 8 words. Other values are not supported.")
	}
	for i := 0; i < N; i++ {
		expandedWords[i] = key[i]
	}

	for i := N; i < 4*R; i++ {
		if i%N == 0 {
			expandedWords[i] = expandedWords[i-N] ^ subWord((expandedWords[i-1]<<8)|(expandedWords[i-1]>>24)) ^ RCON[i/N-1]
		} else if N == 8 && i%N == 4 {
			expandedWords[i] = expandedWords[i-N] ^ subWord(expandedWords[i-1])
		} else {
			expandedWords[i] = expandedWords[i-N] ^ expandedWords[i-1]
		}
	}

	return expandedWords
}
