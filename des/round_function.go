package des

// This file defines the round function f and its components

// expand function takes 32-bit input and expands it to 48-bits
func expand(input [4]byte) (output [6]byte) {
	var temp byte = 0
	for ind, val := range EXPANSION {
		if val%8 != 0 {
			temp = (temp << 1) | ((input[(val-1)/8] >> (8 - val%8)) & 1)
		} else {
			temp = (temp << 1) | (input[(val-1)/8] & 1)
		}
		output[ind/8] = temp
	}
	return output
}

// get_bit_at_position takes a 48-bit byte array and returns the bit at position pos (0-indexed)
func get_bit_at_position(input [6]byte, pos int) byte {
	// pos was 0-indexed we convert it to 1-indexed
	pos++
	if pos%8 == 0 {
		return input[(pos-1)/8] & 1
	}
	return (input[(pos-1)/8] >> (8 - pos%8)) & 1
}

// substituteFromSBox performs the substitution from S-Boxes
func substituteFromSBox(input [6]byte) (output [4]byte) {
	var temp byte = 0
	for i := 0; i < 48; i += 6 {
		temp = (temp << 4) | S_BOXES[i/6][(get_bit_at_position(input, i)<<1)|
			get_bit_at_position(input, i+5)][get_bit_at_position(input, i+1)<<3|
			get_bit_at_position(input, i+2)<<2|
			get_bit_at_position(input, i+3)<<1|
			get_bit_at_position(input, i+4)]
		output[i/12] = temp
	}
	return output
}

// permutation performs the final permutation of the round function f
func permutation(input [4]byte) (output [4]byte) {
	var temp byte = 0
	for ind, val := range PERMUTATION {
		if val%8 != 0 {
			temp = (temp << 1) | ((input[(val-1)/8] >> (8 - val%8)) & 1)
		} else {
			temp = (temp << 1) | (input[(val-1)/8] & 1)
		}
		output[ind/8] = temp
	}
	return output
}

func round_function(input [4]byte, roundkey [6]byte) [4]byte {
	xored_input := [6]byte{}
	expanded := expand(input)
	for i := 0; i < 6; i++ {
		xored_input[i] = expanded[i] ^ roundkey[i]
	}
	return permutation(substituteFromSBox(xored_input))
}
