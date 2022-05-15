package isogram

import "strings"

func IsIsogram(word string) bool {

	for i := 0; i < len(word); i++ {
		for x := i + 1; x < len(word); x++ {
			if strings.ToLower(string(word[i])) != " " &&
				strings.ToLower(string(word[i])) != "-" &&
				strings.ToLower(string(word[i])) == strings.ToLower(string(word[x])) {
				return false
			}
		}
	}
	return true
}
