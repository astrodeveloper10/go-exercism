package raindrops

import "strconv"

func Convert(number int) string {
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0{
		return strconv.Itoa(number)
    }

    message := ""

    if number % 3 == 0 {
        message += "Pling"
    }
    if number % 5 == 0 {
        message += "Plang"
    }
    if number % 7 == 0 {
        message += "Plong"
    }

    return message
}
