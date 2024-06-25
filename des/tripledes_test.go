package des

import "testing"

func TestTripleDES128_Encrypt64(t *testing.T) {
	testKey := []byte("1234567812345678")
	testData := [8]byte([]byte("12345678"))
	expected := [8]byte{0x96, 0xd0, 0x02, 0x88, 0x78, 0xd5, 0x8c, 0x89}
	tdes := new(TripleDES128)
	tdes.Init(testKey)
	encrypted := tdes.Encrypt64(testData)
	if encrypted != expected {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, testData, testKey, encrypted, expected)
	}
}

func TestTripleDES128_Decrypt64(t *testing.T) {
	testKey := []byte("1234567812345678")
	testData := [8]byte{0x96, 0xd0, 0x02, 0x88, 0x78, 0xd5, 0x8c, 0x89}
	expected := [8]byte([]byte("12345678"))
	tdes := new(TripleDES128)
	tdes.Init(testKey)
	decrypted := tdes.Decrypt64(testData)
	if decrypted != expected {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, testData, testKey, decrypted, expected)
	}
}

func TestTripleDES192_Encrypt64(t *testing.T) {
	testKey := []byte("123456781234567887654321")
	testData := [8]byte([]byte("12345678"))
	expected := [8]byte{0x0d, 0xa0, 0x61, 0x56, 0xd0, 0x95, 0x94, 0xc3}
	tdes := new(TripleDES192)
	tdes.Init(testKey)
	encrypted := tdes.Encrypt64(testData)
	if encrypted != expected {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, testData, testKey, encrypted, expected)
	}
}

func TestTripleDES192_Decrypt64(t *testing.T) {
	testKey := []byte("123456781234567887654321")
	testData := [8]byte{0x0d, 0xa0, 0x61, 0x56, 0xd0, 0x95, 0x94, 0xc3}
	expected := [8]byte([]byte("12345678"))
	tdes := new(TripleDES192)
	tdes.Init(testKey)
	decrypted := tdes.Decrypt64(testData)
	if decrypted != expected {
		t.Fatalf(`Testcase: data=%x;key=%x, result: %x, expected: %x`, testData, testKey, decrypted, expected)
	}
}
