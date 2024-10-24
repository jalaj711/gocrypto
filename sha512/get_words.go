package sha512

func getWords(block []byte) [80]uint64 {
	if len(block) != 128 {
		panic("invalid length of block")
	}
	var words [80]uint64

	var i int = 0
	var j int = 0
	var k int = 0
	for i < 16 {
		j = 0
		words[i] = 0
		// generate uint64 from 8 bytes of block
		for j < 8 {
			words[i] = words[i]<<8 | uint64(block[k])
			k++
			j++
		}
		i++
	}

	for i < 80 {
		words[i] = sigma1(words[i-2]) + words[i-7] + sigma0(words[i-15]) + words[i-16]
		i++
	}

	return words
}
