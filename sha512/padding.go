package sha512

import "encoding/binary"

func Pad(input []byte) []byte {
	l := uint64(len(input))
	toPad := l % 128
	if toPad >= 112 {
		toPad = 112 + 128 - toPad
	} else {
		toPad = 112 - toPad
	}
	input = append(input, 0b10000000)
	toPad--
	toPad += 8
	for toPad > 0 {
		input = append(input, 0)
		toPad--
	}
	buf := make([]byte, 8)
	// we want the output in big endian format
	binary.BigEndian.PutUint64(buf, l<<3)
	input = append(input, buf...)
	return input
}
