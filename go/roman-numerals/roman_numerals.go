package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	rN := ""
	i := 0
	if input <= 0 || input > 3000 {
		return "", errors.New("invalid entry")
	}

	for input > 0 {
		switch {
		case input/1000 > 0:
			for i < input/1000 {
				rN = rN + "M"
				i++
			}
			i = 0
			input = input - (1000 * (input / 1000))
		case input >= 900:
			rN = rN + "CM"
			input = input - 900
		case input >= 500:
			rN = rN + "D"
			input = input - 500
		case input >= 400:
			rN = rN + "CD"
			input = input - 400
		case input/100 > 0 && input/100 < 4:
			for i < input/100 {
				rN = rN + "C"
				i++
			}
			i = 0
			input = input - (100 * (input / 100))
		case input >= 90:
			rN = rN + "XC"
			input = input - 90
		case input >= 50:
			rN = rN + "L"
			input = input - 50
		case input >= 40:
			rN = rN + "XL"
			input = input - 40
		case input/10 > 0 && input/10 < 4:
			for i < input/10 {
				rN = rN + "X"
				i++
			}
			i = 0
			input = input - (10 * (input / 10))
		case input >= 9:
			rN = rN + "IX"
			input = input - 9
		case input >= 5:
			rN = rN + "V"
			input = input - 5
		case input >= 4:
			rN = rN + "IV"
			input = input - 4
		case input/1 > 0 && input/1 < 4:
			for i < input/1 {
				rN = rN + "I"
				i++
			}
			input = input - (1 * (input / 1))
		}
	}
	return rN, nil
}
