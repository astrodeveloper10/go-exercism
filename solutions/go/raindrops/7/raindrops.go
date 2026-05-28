package raindrops

import (
	"strconv"
	"strings"
)

var data = []struct {
	divisor int
	msg     string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var sb strings.Builder

	for _, val := range data {
		if number%val.divisor == 0 {
			sb.WriteString(val.msg)
		}
	}

	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(number))
	}

	return sb.String()
}
