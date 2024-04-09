package aes

// helper function to transpose a matrix
// takes a slice of uint32 and treats each uint32 as array of 4 bytes
// this array of 4 byte arrays is transposed and then the result is
// converted to matrix of uint32 again.
//
// Required because in AES a word is taken column wise not row wise
func _transpose(matrix []uint32) []uint32 {
	s := [4][4]uint32{}
	ss := [4][4]uint32{}

	for i := 0; i < 4; i++ {
		s[i][0] = matrix[i] >> 24
		s[i][1] = (matrix[i] >> 16) & 255
		s[i][2] = (matrix[i] >> 8) & 255
		s[i][3] = (matrix[i]) & 255
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			ss[i][j] = s[j][i]
		}
	}
	for i := 0; i < 4; i++ {
		matrix[i] = ss[i][0]<<24 | ss[i][1]<<16 | ss[i][2]<<8 | ss[i][3]
	}
	return matrix
}

func Encrypt128(data []uint32, key []uint32) []uint32 {
	N := 10
	expanded := expandKey(key)
	state := _transpose(data)
	state = addRoundKey(state, _transpose(expanded[0:4]))
	for i := 0; i < N-1; i++ {
		state = subBytes(state)
		state = shiftRows(state)
		state = mixColumns(state)
		state = addRoundKey(state, _transpose(expanded[(i+1)*4:(i+2)*4]))
	}
	state = subBytes(state)
	state = shiftRows(state)
	state = addRoundKey(state, _transpose(expanded[N*4:(N+1)*4]))
	return _transpose(state)
}
