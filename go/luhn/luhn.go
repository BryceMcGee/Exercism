package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var sum int64 = 0

	//stripping spaces
	stripped := strings.ReplaceAll(id, " ", "")

	//reversing the order
	chars := []rune(stripped)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	reversedId := string(result)

	// Check to see if the input is a number
	_, err := strconv.Atoi(reversedId)
	if err != nil {
		return false
	}

	if len(stripped) > 1 {
		for i := 1; i <= len(reversedId); i++ {
			str, _ := strconv.ParseInt(string(reversedId[i-1]), 0, 32)
			if i%2 == 0 {
				str = str * 2
				if str > 9 {
					str = str - 9
				}
			}
			sum += str
		}
		if sum%10 == 0 {
			return true
		}
	}
	return false
}
