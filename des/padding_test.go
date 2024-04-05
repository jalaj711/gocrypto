package des

import (
	"testing"
)

func TestAddPadding(t *testing.T) {
	test := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b}
	expected := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x02, 0x02}
	key := addPadding(test)

	if len(key) != len(expected) {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
	for i := range expected {
		if key[i] != expected[i] {
			t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
		}
	}
}

func TestAddPaddingWholeBlock(t *testing.T) {
	test := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x2, 0x2}
	expected := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x02, 0x02, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08}
	key := addPadding(test)

	if len(key) != len(expected) {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
	for i := range expected {
		if key[i] != expected[i] {
			t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
		}
	}
}

func TestRemovePadding(t *testing.T) {
	expected := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b}
	test := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x02, 0x02}
	key, _ := removePadding(test)

	if len(key) != len(expected) {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
	for i := range expected {
		if key[i] != expected[i] {
			t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
		}
	}
}

func TestRemovePaddingWholeBlock(t *testing.T) {
	expected := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x2, 0x2}
	test := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x02, 0x02, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08}
	key, _ := removePadding(test)

	if len(key) != len(expected) {
		t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
	}
	for i := range expected {
		if key[i] != expected[i] {
			t.Fatalf(`Testcase: %b, result: %b, expected: %b`, test, key, expected)
		}
	}
}

func TestRemovePaddingInvalidPadding(t *testing.T) {
	test := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x02, 0x02, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10}
	_, err := removePadding(test)

	if err == nil {
		t.Fatalf(`Testcase: %b, expected error`, test)
	}
}

func TestRemovePaddingInvalidPadding2(t *testing.T) {
	test := []byte{0x45, 0x67, 0x7a, 0x89, 0x90, 0x9b, 0x02, 0x02, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x0}
	_, err := removePadding(test)

	if err == nil {
		t.Fatalf(`Testcase: %b, expected error`, test)
	}
}
