package product

import (
	"testing"
)

func TestSpeed(t *testing.T) {
	for i := 0; i < 10000; i++ {
		HexToC22("005056B31E4F1EE48398F7FA866FCB15")
		C22ToHex("051MinvF7kI3cFVwXc}B5G")
	}

}

func TestAssertExamples(t *testing.T) {
	data := map[string]string{
		"005056B31E4F1EE48398F7FA866FCB15": "051MinvF7kI3cFVwXc}B5G",
		"005056B33BF11EE3BAC398355F449E72": "051Mipln7kEwmvWrNqIUSW",
		"005056B3483B1ED487A6145349669844": "051MiqWx7jI7fXHJIMQOH0",
		"4C419E52B069732EE10000000A141F1F": "J46UKh1fSoxX00002XGV7m",
		"91AD714ECF7C7571E10000000A141F31": "aQrnJizyTN7X00002XGVCG",
	}

	for k, v := range data {
		if HexToC22(k) != v {
			t.Fatal("Failed HexToC22")
		}
		if C22ToHex(v) != k {
			t.Fatal("Failed C22ToHex")
		}

		if C22ToHex(HexToC22(k)) != k {
			t.Fatal("Failed C22toHex(HexToC22)")
		}

		if HexToC22(C22ToHex(v)) != v {
			t.Fatal("Failed HextoC22(C22ToHex)")
		}

	}
}

func TestSpeed2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		Encode("005056B31E4F1EE48398F7FA866FCB15")
		Decode("051MinvF7kI3cFVwXc}B5G")
	}

}

func TestEncoder(t *testing.T) {

	data := map[string]string{
		"005056B31E4F1EE48398F7FA866FCB15": "051MinvF7kI3cFVwXc}B5G",
		"005056B33BF11EE3BAC398355F449E72": "051Mipln7kEwmvWrNqIUSW",
		"005056B3483B1ED487A6145349669844": "051MiqWx7jI7fXHJIMQOH0",
		"4C419E52B069732EE10000000A141F1F": "J46UKh1fSoxX00002XGV7m",
		"91AD714ECF7C7571E10000000A141F31": "aQrnJizyTN7X00002XGVCG",
	}

	for k, v := range data {
		if Encode(k) != v {
			t.Fatal("Failed HexToC22")
		}
		if Decode(v) != k {
			t.Fatal("Failed C22ToHex")
		}

		if Decode(Encode(k)) != k {
			t.Fatal("Failed C22toHex(HexToC22)")
		}

		if Encode(Decode(v)) != v {
			t.Fatal("Failed HextoC22(C22ToHex)")
		}

	}

}
