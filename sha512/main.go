package sha512

type registers struct {
	a uint64
	b uint64
	c uint64
	d uint64
	e uint64
	f uint64
	g uint64
	h uint64
}

func Hash(input []byte) []byte {
	input = Pad(input)

	reg := getInitRegisters()

	for i := 0; i < len(input); i += 128 {
		sha512block(input[i:i+128], &reg)
	}

	result := []uint64{reg.a, reg.b, reg.c, reg.d, reg.e, reg.f, reg.g, reg.h}

	return uint64ToByte(result)
}
