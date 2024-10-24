package sha512

// the round function of SHA512
// W: The word for this round
// H: Pointer to the registers struct. THis is updated in place
// t: 0-indexed round number
func roundF(W uint64, H *registers, t int) {
	var t1, t2 uint64
	t1 = H.h + ch(H.e, H.f, H.g) + bigSigma1(H.e) + W + K[t]
	t2 = bigSigma0(H.a) + maj(H.a, H.b, H.c)
	H.h = H.g
	H.g = H.f
	H.f = H.e
	H.e = H.d + t1
	H.d = H.c
	H.c = H.b
	H.b = H.a
	H.a = t1 + t2
}
