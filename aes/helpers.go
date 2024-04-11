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

// helper function to perform multiplication in GF(256)
func _multiply(num byte, multiplyBy byte) byte {
	if multiplyBy&1 == 1 {
		return num ^ _multiply(num, multiplyBy^1)
	}
	ans := num << 1
	if num&0x80 == 0x80 {
		ans ^= 0x1b
	}
	multiplyBy = multiplyBy >> 1
	if multiplyBy > 1 {
		return _multiply(ans, multiplyBy)
	}
	return ans
}

// converts byte arrays to uint32 arrays
func _byteToUintArr(arr []byte) []uint32 {
	if len(arr)%4 != 0 {
		panic("Cannot convert byte array of invalid length to uint32 array")
	}
	converted := make([]uint32, len(arr)/4)
	for i := 0; i < len(arr); i += 4 {
		converted[i/4] = uint32(arr[i])<<24 | uint32(arr[i+1])<<16 | uint32(arr[i+2])<<8 | uint32(arr[i+3])
	}

	return converted
}

// converts uint32 arrays to byte arrays
func _uintArrToByte(arr []uint32) []byte {
	converted := make([]byte, len(arr)*4)
	for i := 0; i < len(arr); i++ {
		converted[i*4] = byte(arr[i] >> 24)
		converted[i*4+1] = byte((arr[i] >> 16) & 255)
		converted[i*4+2] = byte((arr[i] >> 8) & 255)
		converted[i*4+3] = byte(arr[i] & 255)
	}
	return converted
}
