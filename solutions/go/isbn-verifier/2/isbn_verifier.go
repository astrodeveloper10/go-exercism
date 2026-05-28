package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
    n := len(isbn)

	if n != 10 {
		return false
	}
	
	sum := 0

	for i, c := range isbn {
        if unicode.IsDigit(c) {
            sum += int(c-'0') * n
        } else if string(c) == "X" && i == 9 {
            sum += 10
        } else {
            return false
        }
        
		n -= 1
	}

	return sum%11 == 0
}
