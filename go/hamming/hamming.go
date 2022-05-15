package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hamCount := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hamCount++
			}
		}
		return hamCount, nil
	}
	return 0, errors.New("string lengths do not match")
}
