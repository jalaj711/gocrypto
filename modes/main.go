package modes

// TODO add a isinitiated function to test if the cipher has been initiated
// TODO add function to get and set expanded key, current expanded key gets manipulated
type BlockCipher interface {
	Init([]byte)
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
	GetBlockSize() int
}
