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

func Convert(number int) string {
	var sb strings.Builder

    for i := 3; i <= 7; i += 2 {
        if number % i == 0 {
            sb.WriteString(msgMap[i])
        }
    }

	if sb.Len() == 0 {
		sb.WriteString(strconv.Itoa(number))
	}

	return sb.String()
}
