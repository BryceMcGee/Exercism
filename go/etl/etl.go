package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transform := make(map[string]int)

	for i := range in {
		for _, v := range in[i] {
			transform[strings.ToLower(v)] = i
		}
	}
	return transform
}
