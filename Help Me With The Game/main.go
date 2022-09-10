package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	charactersPerLine = 34
	asciiLowerToUpperDiff = 32
)

var rank = map[byte]int { 'K': 0, 'Q': 1, 'R': 2, 'B': 3, 'N': 4, 'P': 5 }

type chessNotation struct {
	piece   byte
	column  byte
	row     int
}

func (n chessNotation) String() string {
	if n.piece == 'P' {
		return fmt.Sprintf("%s%d", string(n.column), n.row)
	}
	return fmt.Sprintf("%s%s%d", string(n.piece), string(n.column), n.row)
}

type player struct {
	isWhite        bool
	chessNotations []chessNotation
}

func (p player) sort() {
	sort.Slice(p.chessNotations, func(i, j int) bool {
	// Rank has highest value
	if rank[p.chessNotations[i].piece] != rank[p.chessNotations[j].piece] {
		return rank[p.chessNotations[i].piece] < rank[p.chessNotations[j].piece]
	}
	// Row has second highest value
	if p.chessNotations[i].row != p.chessNotations[j].row {
		// White pieces are listed by lowest->highest row
		if p.isWhite {
			return p.chessNotations[i].row < p.chessNotations[j].row
		}
		// Black pieces are listed by highest->lowest row
		return p.chessNotations[i].row > p.chessNotations[j].row
	}
	// Column has lowest value
	return p.chessNotations[i].column < p.chessNotations[j].column
	})
}

func (p player) String() string {
	strSlice := make([]string, len(p.chessNotations))
	for i, chessNotation := range(p.chessNotations) {
		strSlice[i] = chessNotation.String()
	}
	return strings.Join(strSlice, ",")
}

func main() {
	columnDef := [8]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	white := player { isWhite: true, chessNotations: make([]chessNotation, 0, 16) }
	black := player { isWhite: false, chessNotations: make([]chessNotation, 0, 16) }

	var line string

	for row := 8; row > 0; row-- {
		fmt.Scanln(&line)  // ignore every second line, starting with the first one
		fmt.Scanln(&line)
		for col, charPos := 0, 2; charPos < charactersPerLine; col, charPos = col+1, charPos+4 {
			var char = line[charPos]
			if char >= 'B' && char <= 'R' {
				notation := chessNotation { piece: char, column: columnDef[col], row: row }
				white.chessNotations = append(white.chessNotations, notation)
			}
			if char >= 'b' && char <= 'r' {
				notation := chessNotation { piece: toUpper(char), column: columnDef[col], row: row }
				black.chessNotations = append(black.chessNotations, notation)
			}
		}
	}

	white.sort()
	black.sort()
	fmt.Printf("White: %s\n", white)
	fmt.Printf("Black: %s\n", black)
}

func toUpper(lowerCase byte) byte {
	return lowerCase - asciiLowerToUpperDiff
}
