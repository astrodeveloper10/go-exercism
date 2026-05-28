package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
	noSpace := strings.ReplaceAll(id, " ", "")

    if len(noSpace) <= 1 {
        return false
    }
    
	b := []rune(noSpace)
	n := len(b)
	sum := 0

	for _, r := range noSpace {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	for i := n - 2; i >= 0; i -= 2 {
		r := rune(b[i])
		num := int(r-'0') * 2

		if num > 9 {
			num -= 9
		}

		b[i] = rune(num)
	}

	for _, r := range b {
		if r > 9 {
			sum += int(r - '0')
		} else {
			sum += int(r)
		}
	}

	return sum%10 == 0
}
