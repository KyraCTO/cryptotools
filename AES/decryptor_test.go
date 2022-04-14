package AES

import (
	"fmt"
	"testing"
)

func TestDecryptAESPayload(t *testing.T) {
	fmt.Println("Running Test...")
	key1 := GenerateAESKey()
	key2 := GenerateAESKey()

	if key1 == key2 {
		t.Errorf("Generated Keys match. They should not.")
	}
}
