package modes

type Padding interface {
	Pad([]byte, int) []byte
	UnPad([]byte, int) []byte
}

type PKCS7 struct{}

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
