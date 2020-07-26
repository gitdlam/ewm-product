package product

import (
	"bytes"

	//	"log"
	"math/big"
	"strings"
	//"strconv"
)

// This is added later based on the standard library code. Runs 10x faster.
func Encode(hex string) string {
	bigInt := new(big.Int)
	bigInt.SetString(hex, 16)
	return RawCustomEncoding.EncodeToString(append(bytes.Repeat([]byte{byte(0)}, 16-len(bigInt.Bytes())), bigInt.Bytes()...))

}

// This is added later based on the standard library code. Runs 10x faster.
func Decode(c22 string) string {
	b, _ := RawCustomEncoding.DecodeString(c22)
	bigInt := new(big.Int)
	bigInt.SetBytes(b)
	s := strings.ToUpper(bigInt.Text(16))
	return strings.Repeat("0", 32-len(s)) + s
}
