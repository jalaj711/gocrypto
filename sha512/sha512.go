package sha512

func sha512block(block []byte, reg *registers) {
	if len(block) != 128 {
		panic("invalid block length")
	}

	words := getWords(block)
	regCopy := *reg

	for i := 0; i < 80; i++ {
		roundF(words[i], reg, i)
	}

	reg.a += regCopy.a
	reg.b += regCopy.b
	reg.c += regCopy.c
	reg.d += regCopy.d
	reg.e += regCopy.e
	reg.f += regCopy.f
	reg.g += regCopy.g
	reg.h += regCopy.h

}
