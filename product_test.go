package product

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	hexString := "000D3AD2A5EB1ED9A6F84A4E0A7E3DE2"

	fmt.Println("hex:", hexString)
	uuid := HexToUU(hexString)
	fmt.Println("uuid:", uuid)
	fmt.Println("hex:", UUToHex(uuid))
	if UUToHex(uuid) != hexString {
		t.Fatal("Strings should equal.")
	}

}
