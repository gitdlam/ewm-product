package product

import (
	"fmt"

	"math/big"
	"strings"
	//"strconv"
)

var (
	base64Table string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz{}"
	indexMap    map[rune]int
)

func init() {
	indexMap = map[rune]int{}

	for i, ch := range base64Table {
		indexMap[ch] = i
	}
}

func C22ToHex(c22 string) string {
	binaryChunks := make([]string, 0, 22)
	for _, ch := range c22 {
		// each number becomes a chunk of 6 bits.
		binaryChunks = append(binaryChunks, fmt.Sprintf("%06b", indexMap[ch]))

	}
	bigInt := new(big.Int)
	b := strings.Join(binaryChunks, "")
	// UUID is 128 bits.
	// 22 chunks of 6 bits = 22*6 = 132 bits which is longer than 128 bits by 4.  Therefore the last 4 bits are discarded.
	b = b[:len(b)-4]
	bigInt.SetString(b, 2)
	h := bigInt.Text(16)
	// padding zeros to make full 32 characters
	h = strings.Repeat("0", 32-len(h)) + h
	return strings.ToUpper(h)

}

func HexToC22(hex string) string {
	c22 := make([]byte, 0, 22)

	bigInt := new(big.Int)
	bigInt.SetString(hex, 16)
	b := bigInt.Text(2)
	// UUID is 128 bits. Pad zeros on left, and add 4 bits, to make 132 bits, which will become 22 chunks of 6 bits.
	b = strings.Repeat("0", 128-len(b)) + b + "0000"

	//fmt.Println(b)

	i := 0
	j := 6

	for i < 132 {
		bigInt.SetString(b[i:j], 2)
		c22 = append(c22, base64Table[int(bigInt.Int64())])
		i = i + 6
		j = j + 6
	}

	return string(c22)
}
