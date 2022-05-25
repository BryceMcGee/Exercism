package dna

import "errors"

type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	if len(d) >= 0 {
		for _, v := range d {
			switch v {
			case 'A':
				h['A']++
			case 'C':
				h['C']++
			case 'G':
				h['G']++
			case 'T':
				h['T']++
			default:
				return h, errors.New("invalid character")
			}
		}
	}
	return h, nil
}
