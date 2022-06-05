package protein

import (
	"errors"
	"strings"
)

var ErrStop = errors.New("err stop")
var ErrInvalidBase = errors.New("err invalid")

var m = make(map[string]string)

//init sets up the translation map from codons to its corresponding proteins.
func init() {
	m["AUG"] = "Methionine"
	m["UUU"] = "Phenylalanine"
	m["UUC"] = "Phenylalanine"
	m["UUA"] = "Leucine"
	m["UUG"] = "Leucine"
	m["UCU"] = "Serine"
	m["UCC"] = "Serine"
	m["UCA"] = "Serine"
	m["UCG"] = "Serine"
	m["UAU"] = "Tyrosine"
	m["UAC"] = "Tyrosine"
	m["UGU"] = "Cysteine"
	m["UGC"] = "Cysteine"
	m["UGG"] = "Tryptophan"
	m["UAA"] = "STOP"
	m["UAG"] = "STOP"
	m["UGA"] = "STOP"
}

//FromRNA takes in a given string (example "AUGUUUUCUUAAAUG") and returns the slice of matching proteins.
func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	newRna := ConvertString(rna)
	var s []string

	for i := range newRna {
		if _, ok := m[newRna[i]]; ok && m[newRna[i]] == "STOP" {
			break
		} else {
			s = append(s, m[newRna[i]])
		}
	}
	return s, nil
}

//FromCodon takes in a given string (example: "UCU", "UCC", "UCA", or "UCG") and returns the matching protein.
func FromCodon(codon string) (string, error) {
	if val, ok := m[codon]; ok && val != "STOP" {
		return val, nil
	} else if val == "STOP" {
		return "", ErrStop
	} else {
		return "", ErrInvalidBase
	}
}

//ConvertString takes in a given string, adds a space every third character, splits and returns the group into a slice.
func ConvertString(rna string) []string {
	newRna := ""
	for i := 0; i < len(rna); i++ {
		newRna += string(rna[i])
		if (i%3)-2 == 0 {
			newRna += " "
		}
	}
	s := strings.Split(newRna, " ")
	s = s[:len(s)-1]

	return s
}
