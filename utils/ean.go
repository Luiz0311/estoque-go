package utils

import (
	"math/rand/v2"
	"strconv"
)

func GenerateEAN13() string {
	base := ""
	for i := 0; i < 12; i++ {
		base += strconv.Itoa(rand.IntN(10))
	}

	checkSum := calcVeriferDigit(base)

	return base + strconv.Itoa(checkSum)
}

func calcVeriferDigit(code string) int {
	sum := 0
	for i, r := range code {
		digit := int(r - '0')
		if (i+1)%2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	return (10 - (sum % 10)) % 10
}
