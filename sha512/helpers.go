package sha512

// circular right shift the number x by n bits
func rotr(x uint64, n int) uint64 {
	return x>>n | (x << (64 - n))
}

// right shift the number x by n bits
func shr(x uint64, n int) uint64 {
	return x >> n
}

func bigSigma0(a uint64) uint64 {
	return rotr(a, 28) ^ rotr(a, 34) ^ rotr(a, 39)
}

func bigSigma1(a uint64) uint64 {
	return rotr(a, 14) ^ rotr(a, 18) ^ rotr(a, 41)
}

func sigma0(a uint64) uint64 {
	return rotr(a, 1) ^ rotr(a, 8) ^ shr(a, 7)
}

func sigma1(a uint64) uint64 {
	return rotr(a, 19) ^ rotr(a, 61) ^ shr(a, 6)
}

// condition function: if e then f else g
// used in round function
func ch(e, f, g uint64) uint64 {
	return (e & f) ^ (^e & g)
}

// majority function: true when at least 2 of the 3 params are true
// used in round function
func maj(a, b, c uint64) uint64 {
	return (a & b) ^ (a & c) ^ (b & c)
}

// convert uint64 array to byte array
func uint64ToByte(inp []uint64) []byte {
	n := len(inp)
	out := make([]byte, n*8)

	for i := 0; i < n; i++ {
		v := inp[i]
		for j := 0; j < 8; j++ {
			out[i*8+7-j] = byte(v & 0xff)
			v = v >> 8
		}
	}

	return out
}
