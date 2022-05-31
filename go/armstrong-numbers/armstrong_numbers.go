package armstrong

import "strconv"

func IsNumber(n int) bool {

	num := ConvertNumber(n)
	if num == n {
		return true
	}
	return false
}

func ConvertNumber(n int) int {
	fNum := 0
	s := strconv.Itoa(n)
	digits := len(s)
	for i := range s {
		fNum += Powers(int(s[i]-48), digits)
	}
	return fNum
}

func Powers(number int, digits int) int {

	if digits == 1 {
		return number
	}

	power := number

	for i := 1; i < digits; i++ {
		number = number * power
	}
	return number
}
