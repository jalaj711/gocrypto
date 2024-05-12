package aes

type AES struct {
	__expanded []uint32
	Nr         int
}

func (aes *AES) GetBlockSize() int {
	return 16
}

// Init : Initialise the AES object with the supplied key
func (aes *AES) Init(key []byte) {
	if len(key) == 16 {
		aes.Nr = 10
	} else if len(key) == 24 {
		aes.Nr = 12
	} else if len(key) == 32 {
		aes.Nr = 14
	} else {
		panic("Invalid key length, should be either 128, 192 or 256 bits long")
	}
	aes.__expanded = expandKey(_byteToUintArr(key))
}

// Encrypt128 : Encrypts a single block of 128 bits passed as an array
// of 4 32-bit integers
func (aes *AES) Encrypt128(data []uint32) []uint32 {
	_transpose(data)
	// TODO remove this
	expandedCopy := make([]uint32, len(aes.__expanded))
	copy(expandedCopy, aes.__expanded)
	addRoundKey(data, _transpose(expandedCopy[0:4]))
	for i := 0; i < aes.Nr-1; i++ {
		subBytes(data)
		shiftRows(data)
		mixColumns(data)
		addRoundKey(data, _transpose(expandedCopy[(i+1)*4:(i+2)*4]))
	}
	subBytes(data)
	shiftRows(data)
	addRoundKey(data, _transpose(expandedCopy[aes.Nr*4:(aes.Nr+1)*4]))
	return _transpose(data)
}

// Encrypt : Encrypts a single block of 128 bits passes as an array
// of 16 bytes
func (aes *AES) Encrypt(data []byte) []byte {
	return _uintArrToByte(aes.Encrypt128(_byteToUintArr(data)))
}

// Decrypt128 : Decrypts a single block of 128 bits passed as an array
// of 4 32-bit integers
func (aes *AES) Decrypt128(data []uint32) []uint32 {
	_transpose(data)
	// TODO remove this
	expandedCopy := make([]uint32, len(aes.__expanded))
	copy(expandedCopy, aes.__expanded)
	addRoundKey(data, _transpose(expandedCopy[aes.Nr*4:(aes.Nr+1)*4]))
	for i := aes.Nr - 1; i > 0; i-- {
		invShiftRows(data)
		invSubBytes(data)
		addRoundKey(data, _transpose(expandedCopy[(i)*4:(i+1)*4]))
		invMixColumns(data)
	}
	invShiftRows(data)
	invSubBytes(data)
	addRoundKey(data, _transpose(expandedCopy[0:4]))
	return _transpose(data)
}

// Decrypt : Decrypts a single block of 128 bits passes as an array
// of 16 bytes
func (aes *AES) Decrypt(data []byte) []byte {
	return _uintArrToByte(aes.Decrypt128(_byteToUintArr(data)))
}
