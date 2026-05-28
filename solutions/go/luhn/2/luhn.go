package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	n := len(id)

	if n <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := n - 1; i >= 0; i-- {
		r := rune(id[i])
		if !unicode.IsDigit(r) {
			return false
		}

		num := int(r - '0')

		if double {
			num *= 2

			if num > 9 {
				num -= 9
			}
		}

		sum += num
		double = !double

	}

	return sum%10 == 0
}
