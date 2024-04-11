package aes

type AES struct {
	__expanded []uint32
	Nr         int
}

type _AES interface {
	Init([]byte)
	Encrypt([]byte) []byte
}

func (aes *AES) Init(key []byte) {
	if len(key) == 8 {
		aes.Nr = 10
	} else if len(key) == 12 {
		aes.Nr = 12
	} else if len(key) == 16 {
		aes.Nr = 16
	} else {
		panic("Invalid key length, should be either 128, 192 or 256 bits long")
	}
	aes.__expanded = expandKey(_byteToUintArr(key))
}

// Encrypt128 : Encrypts a single block of 128 bits passed as an array
// of 4 32-bit integers
func (aes *AES) Encrypt128(data []uint32) []uint32 {
	_transpose(data)
	addRoundKey(data, _transpose(aes.__expanded[0:4]))
	for i := 0; i < aes.Nr-1; i++ {
		subBytes(data)
		shiftRows(data)
		mixColumns(data)
		addRoundKey(data, _transpose(aes.__expanded[(i+1)*4:(i+2)*4]))
	}
	subBytes(data)
	shiftRows(data)
	addRoundKey(data, _transpose(aes.__expanded[aes.Nr*4:(aes.Nr+1)*4]))
	return _transpose(data)
}

// Encrypt : Encrypts a single block of 128 bits passes as an array
// of 16 bytes
func (aes *AES) Encrypt(data []byte) []byte {
	return _uintArrToByte(aes.Encrypt128(_byteToUintArr(data)))
}
