package scrabble

func Score(word string) int {
	points := 0

	for i := 0; i < len(word); i++ {
		points += Convert(string(word[i]))
	}
	return points
}

func Convert(letter string) int {
	switch letter {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U", "l", "L", "n", "N", "r", "R", "s", "S", "t", "T":
		return 1
	case "d", "D", "g", "G":
		return 2
	case "b", "B", "c", "C", "m", "M", "p", "P":
		return 3
	case "f", "F", "h", "H", "v", "V", "w", "W", "y", "Y":
		return 4
	case "k", "K":
		return 5
	case "j", "J", "x", "X":
		return 8
	case "q", "Q", "z", "Z":
		return 10
	default:
		return 0
	}
}

/*
Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
