package product

import (
	"fmt"

	"math/big"
	"strings"
	//"strconv"
)

var (
	schemeString string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz{}"
	scheme       map[string]int
)

func init() {
	scheme = map[string]int{}

	for i, ch := range schemeString {
		scheme[string(ch)] = i
	}
}

func UUToHex(uuid string) string {
	uuidSlice := make([]string, 0, 22)
	for _, ch := range uuid {
		uuidSlice = append(uuidSlice, fmt.Sprintf("%06b", scheme[string(ch)]))

	}
	bigInt := new(big.Int)
	b := strings.Join(uuidSlice, "")
	b = b[:len(b)-4]
	bigInt.SetString(b, 2)
	h := bigInt.Text(16)
	h = strings.Repeat("0", 32-len(h)) + h
	return strings.ToUpper(h)

}

func HexToUU(s string) string {
	uuidSlice := make([]string, 0, 22)

	bigInt := new(big.Int)
	bigInt.SetString(s, 16)
	b := bigInt.Text(2)

	b = strings.Repeat("0", 128-len(b)) + b + "0000"

	//fmt.Println(b)

	i := 0
	j := 6

	for i < 132 {
		bigInt.SetString(b[i:j], 2)
		uuidSlice = append(uuidSlice, schemeString[int(bigInt.Int64()):int(bigInt.Int64())+1])
		i = i + 6
		j = j + 6
	}

	return strings.Join(uuidSlice, "")
}
