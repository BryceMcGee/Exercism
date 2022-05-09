package chessboard

// Rank Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	occupied := 0
	row, rnk := cb[rank]

	if !rnk {
		return 0
	}
	for _, v := range row {
		if v {
			occupied++
		}
	}
	return occupied
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	occupied := 0

	if file < 1 || file > 8 {
		return 0
	}
	for _, v := range cb {
		if v[file-1] {
			occupied++
		}
	}
	return occupied
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	totalSquares := 0

	for _, v := range cb {
		totalSquares += len(v)
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	occupied := 0

	for i := range cb {
		occupied += CountInRank(cb, i)
	}
	return occupied
}
