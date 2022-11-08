package chessboard

// File, the verticals of a chess board
type File []bool // rank
// Chessboard, made up of 8 Files, A-H
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var count = 0

	for _, piecePresent := range cb[file] {
		if piecePresent {
			count = count + 1
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank <= 0 || rank > 8 {
		return 0
	}
	var count = 0

	for _, files := range cb {
		if files[rank-1] {
			count = count + 1
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count = 0

	for _, file := range cb {
		count = count + len(file)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count = 0

	for _, file := range cb {
		for _, present := range file {
			if present {
				count = count + 1
			}
		}
	}
	return count
}
