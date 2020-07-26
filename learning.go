package product

import (
	"fmt"

	//	"log"
	"math/big"
	"strings"
	//"strconv"
)

var (
	base64Table string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz{}"
)

func C22ToHex(c22 string) string {
	var binaryChunks strings.Builder
	binaryChunks.Grow(132)
	for _, ch := range c22 {
		// each number becomes a chunk of 6 bits.

		binaryChunks.WriteString(fmt.Sprintf("%06b", decodeRune(ch)))

	}
	bigInt := new(big.Int)

	// UUID is 128 bits.
	// 22 chunks of 6 bits = 22*6 = 132 bits which is longer than 128 bits by 4.  Therefore the last 4 bits are discarded.
	bigInt.SetString(binaryChunks.String()[:binaryChunks.Len()-4], 2)
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

// not signifcantly faster than map, if it is any faster.
func decodeRune(ch rune) byte {
	switch ch {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'A':
		return 10
	case 'B':
		return 11
	case 'C':
		return 12
	case 'D':
		return 13
	case 'E':
		return 14
	case 'F':
		return 15
	case 'G':
		return 16
	case 'H':
		return 17
	case 'I':
		return 18
	case 'J':
		return 19
	case 'K':
		return 20
	case 'L':
		return 21
	case 'M':
		return 22
	case 'N':
		return 23
	case 'O':
		return 24
	case 'P':
		return 25
	case 'Q':
		return 26
	case 'R':
		return 27
	case 'S':
		return 28
	case 'T':
		return 29
	case 'U':
		return 30
	case 'V':
		return 31
	case 'W':
		return 32
	case 'X':
		return 33
	case 'Y':
		return 34
	case 'Z':
		return 35
	case 'a':
		return 36
	case 'b':
		return 37
	case 'c':
		return 38
	case 'd':
		return 39
	case 'e':
		return 40
	case 'f':
		return 41
	case 'g':
		return 42
	case 'h':
		return 43
	case 'i':
		return 44
	case 'j':
		return 45
	case 'k':
		return 46
	case 'l':
		return 47
	case 'm':
		return 48
	case 'n':
		return 49
	case 'o':
		return 50
	case 'p':
		return 51
	case 'q':
		return 52
	case 'r':
		return 53
	case 's':
		return 54
	case 't':
		return 55
	case 'u':
		return 56
	case 'v':
		return 57
	case 'w':
		return 58
	case 'x':
		return 59
	case 'y':
		return 60
	case 'z':
		return 61
	case '{':
		return 62
	case '}':
		return 63
	default:
		return 255
	}
}
