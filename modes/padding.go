package modes

type Padding interface {
	Pad([]byte, int) []byte
	UnPad([]byte, int) []byte
}

type PKCS7 struct{}
type ISO struct{}

func (pk *PKCS7) Pad(data []byte, blockSize int) []byte {
	size := len(data)
	paddingBytes := blockSize - size%blockSize
	padding := make([]byte, paddingBytes)
	for i := range padding {
		padding[i] = byte(paddingBytes)
	}
	return append(data, padding...)
}

func (pk *PKCS7) UnPad(data []byte, blockSize int) []byte {
	if data[len(data)-1] > byte(blockSize) || data[len(data)-1] < 1 {
		panic("invalid padding in data")
	}
	return data[:len(data)-int(data[len(data)-1])]
}

func (iso *ISO) Pad(data []byte, blockSize int) []byte {
	size := len(data)
	paddingBytes := blockSize - size%blockSize
	if paddingBytes == 0 {
		paddingBytes = 1
	}
	padding := make([]byte, paddingBytes)
	padding[0] = 0x80
	for i := 1; i < paddingBytes; i++ {
		padding[i] = 0x0
	}
	return append(data, padding...)
}

func (iso *ISO) UnPad(data []byte, blockSize int) []byte {
	if data[len(data)-1] != 0x0 && data[len(data)-1] != 0x80 {
		panic("invalid padding in data")
	}
	i := len(data) - 1
	for ; i >= 0 && data[i] != 0x80; i-- {
	}
	return data[:i]
}
