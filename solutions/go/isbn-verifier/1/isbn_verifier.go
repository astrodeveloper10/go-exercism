package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}
	
	n := 10
	sum := 0

	for i, c := range isbn {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
            if string(c) == "X" {
                if i != 9 {
                    return false
                }
                sum += 10 * n
            } else {
				sum += int(c-'0') * n
            }
		}

		n -= 1
	}

	return sum%11 == 0
}
