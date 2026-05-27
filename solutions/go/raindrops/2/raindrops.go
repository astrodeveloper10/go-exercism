package raindrops

import (
    "strconv"
    "strings"
    )

func addSound(sb *strings.Builder, number int, divisor int, sound string) {
	if number%divisor == 0 {
		sb.WriteString(sound)
	}
}

func Convert(number int) string {
	var sb strings.Builder

	addSound(&sb, number, 3, "Pling")
	addSound(&sb, number, 5, "Plang")
	addSound(&sb, number, 7, "Plong")

	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(number))
	}

	return sb.String()
}
