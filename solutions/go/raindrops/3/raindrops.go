package raindrops

import (
    "strconv"
    "strings"
    )

var msgMap = map[int]string{
    3 : "Pling",
    5 : "Plang",
    7 : "Plong",
}

func addSound(sb *strings.Builder, number int, divisor int) {
	if number%divisor == 0 {
		sb.WriteString(msgMap[divisor])
	}
}

func Convert(number int) string {
	var sb strings.Builder

    for i := 3; i <= 7; i += 2 {
        addSound(&sb, number, i)
    }

	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(number))
	}

	return sb.String()
}
