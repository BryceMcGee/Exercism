package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	var converted strings.Builder
	if number%3 == 0 {
		converted.WriteString("Pling")
	}
	if number%5 == 0 {
		converted.WriteString("Plang")
	}
	if number%7 == 0 {
		converted.WriteString("Plong")
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return fmt.Sprint(number)
	}
	return converted.String()
}
