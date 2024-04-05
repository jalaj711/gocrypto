package des

import (
	"testing"
)

func TestExpansionFunction(t *testing.T) {
	test := [4]byte{0b11110000, 0b10101010, 0b11110000, 0b10101010}
	expected := [6]byte{0b01111010, 0b00010101, 0b01010101, 0b01111010, 0b00010101, 0b01010101}
	key := expand(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
}

func TestSBoxSubstitution(t *testing.T) {
	test := [6]byte{0b01100001, 0b00010111, 0b10111010, 0b10000110, 0b01100101, 0b00100111}
	expected := [4]byte{0b01011100, 0b10000010, 0b10110101, 0b10010111}
	key := substituteFromSBox(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
}

func TestPermutation(t *testing.T) {
	test := [4]byte{0b01011100, 0b10000010, 0b10110101, 0b10010111}
	expected := [4]byte{0b00100011, 0b01001010, 0b10101001, 0b10111011}
	key := permutation(test)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if key != expected {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
}

func TestRoundFunction(t *testing.T) {
	test_data := [4]byte{0b11110000, 0b10101010, 0b11110000, 0b10101010}
	test_key := [6]byte{0b00011011, 0b00000010, 0b11101111, 0b11111100, 0b01110000, 0b01110010}
	expected := [4]byte{0b00100011, 0b01001010, 0b10101001, 0b10111011}
	out := round_function(test_data, test_key)
	// t.Logf(`Testcase: %b, result: %d`, test, key)
	if out != expected {
		t.Fatalf(`Testcase: data=%b;key=%b, result: %b, expected: %b`, test_data, test_key, out, expected)
	}
}
