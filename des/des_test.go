package des

import (
	"testing"
)

func TestEncrypt64(t *testing.T) {
	testKey := []byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	testData := [8]byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	expected := [8]byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	des := new(DES)
	des.Init(testKey)
	encrypted := des.Encrypt64(testData)
	if encrypted != expected {
		t.Fatalf(`Testcase: data=%b;key=%b, result: %b, expected: %b`, testData, testKey, encrypted, expected)
	}
}

func TestDecrypt64(t *testing.T) {
	testKey := []byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	testData := [8]byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	expected := [8]byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	des := new(DES)
	des.Init(testKey)
	decrypted := des.Decrypt64(testData)
	if decrypted != expected {
		t.Fatalf(`Testcase: data=%b;key=%b, result: %b, expected: %b`, testData, testKey, decrypted, expected)
	}
}

func TestEncrypt(t *testing.T) {
	testKey := []byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	testData := []byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	expected := []byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	des := new(DES)
	des.Init(testKey)
	encrypted := des.Encrypt(testData)
	if len(encrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, testData, testKey, encrypted, expected)
	}
	for i := range expected {
		if encrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, testData, testKey, encrypted, expected, i)
		}
	}
}

func TestDecrypt(t *testing.T) {
	testKey := []byte{0b00010011, 0b00110100, 0b01010111, 0b01111001, 0b10011011, 0b10111100, 0b11011111, 0b11110001}
	testData := []byte{0b10000101, 0b11101000, 0b00010011, 0b01010100, 0b00001111, 0b00001010, 0b10110100, 0b00000101}
	expected := []byte{0b00000001, 0b00100011, 0b01000101, 0b01100111, 0b10001001, 0b10101011, 0b11001101, 0b11101111}
	des := new(DES)
	des.Init(testKey)
	decrypted := des.Decrypt(testData)
	if len(decrypted) != len(expected) {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, testData, testKey, decrypted, expected)
	}
	for i := range expected {
		if decrypted[i] != expected[i] {
			t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x, diff=%d`, testData, testKey, decrypted, expected, i)
		}
	}
}

func TestDecryptInvalidLength(t *testing.T) {
	defer func() {
		//empty function so that program does not exit when panic is raised
		if recover() == nil {
			t.Fatalf(`Testcase: expected error`)
		}
	}()

	testKey := []byte("64bitKey")
	testData := []byte{0x53, 0x62, 0x38, 0x3a, 0x75, 0x2c, 0xf4}
	des := new(DES)
	des.Init(testKey)
	des.Decrypt(testData)
}
