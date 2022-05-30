package wordcount

import (
	"log"
	"regexp"
	"strings"
)

type Frequency map[string]int

//WordCount counts the frequency of a given word in the list that listBuild created
func WordCount(phrase string) Frequency {
	var f = make(Frequency)
	words := parseWords(phrase)
	listBuild(f, words)

	for _, i := range words {
		_, matched := f[i]
		if matched {
			f[i] += 1
		} else {
			f[i] = 1
		}
	}

	return f
}

//parseWords alters the string to remove special/escaped characters and leading/trailing quotes
func parseWords(phrase string) []string {
	splitComma := strings.Split(phrase, ",")
	phrase = strings.Join(splitComma, " ")

	reg, err := regexp.Compile("[^a-zA-Z' \\d]+")
	if err != nil {
		log.Fatal(err)
	}
	regexString := strings.ToLower(reg.ReplaceAllString(phrase, ""))

	splitSpaces := strings.Fields(regexString)

	for i, v := range splitSpaces {
		splitSpaces[i] = strings.Trim(v, "'")
	}

	return splitSpaces
}

//listBuild makes the map of words from the list of strings that parseWords created.
func listBuild(f Frequency, words []string) Frequency {
	for i := range words {
		f[words[i]] = 0
	}
	return f
}
