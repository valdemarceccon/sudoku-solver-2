package solver

import "errors"

type Sudoku [][]int

type Pendent struct {
	X, Y      int
	Possibles []int
}

func (s Sudoku) RowAt(at int) []int {
	return s[at]
}

func (s Sudoku) ColumnAt(at int) (col []int) {
	col = make([]int, 9)
	for index, row := range s {
		col[index] = row[at]
	}

	return
}

func (s Sudoku) QuadrantAt(row int, column int) (quadrant []int) {
	rowStart := (row / 3) * 3
	colStart := (column / 3) * 3
	quadrant = make([]int, 9)

	for rowIndex := rowStart; rowIndex < rowStart+3; rowIndex++ {
		for colIndex := colStart; colIndex < colStart+3; colIndex++ {
			quadrant[(rowIndex-rowStart)*3+(colIndex-colStart)] = s[rowIndex][colIndex]
		}
	}

	return quadrant
}

func (s Sudoku) CheckValidColumn(column int) error {
	if !isGroupValid(s.ColumnAt(column)) {
		return InvalidPuzzleError
	}
	return nil
}

var InvalidPuzzleError = errors.New("invalid puzzle")

func (s Sudoku) CheckValidRow(row int) error {
	if !isGroupValid(s.RowAt(row)) {
		return InvalidPuzzleError
	}

	return nil
}

func (s Sudoku) CheckValidQuadrant(row int, col int) error {
	if !isGroupValid(s.QuadrantAt(row, col)) {
		return InvalidPuzzleError
	}
	return nil
}

func (s Sudoku) FindAllEmpty() []Pendent {
	var result []Pendent
	for rowNumber, row := range s {
		for colNumber, value := range row {
			if value == 0 {
				result = append(result, s.allPossibilitiesAt(rowNumber, colNumber))
			}
		}
	}
	return result
}

func (s Sudoku) allPossibilitiesAt(rowNumber int, colNumber int) Pendent {
	copyPuzzle := s.Copy()
	result := Pendent{
		X:         rowNumber,
		Y:         colNumber,
		Possibles: []int{},
	}
	for v := 1; v <= 9; v++ {
		copyPuzzle[rowNumber][colNumber] = v
		if copyPuzzle.isPuzzleValid(rowNumber, colNumber) {
			result.Possibles = append(result.Possibles, v)
		}
	}

	return result
}

func (s Sudoku) Copy() Sudoku {
	result := make(Sudoku, 9)
	for e := range s {
		result[e] = make([]int, 9)
		copy(result[e], s[e])
	}

	return result
}

func (s Sudoku) isPuzzleValid(rowNumber, colNumber int) bool {
	return s.CheckValidColumn(colNumber) == nil &&
		s.CheckValidRow(rowNumber) == nil &&
		s.CheckValidQuadrant(rowNumber, colNumber) == nil
}

var ImpossiblePuzzle = errors.New("impossible puzzle")

func (s Sudoku) Solve() error {
	empties := s.FindAllEmpty()

	for _, val := range empties {

		if len(val.Possibles) == 0 {
			return ImpossiblePuzzle
		}

		if len(val.Possibles) == 1 {
			s[val.X][val.Y] = val.Possibles[0]
		}
	}

	empties = s.FindAllEmpty()

	if len(empties) != 0 {
		return s.Solve()
	}
	return nil
}

func isGroupValid(col []int) bool {
	numbers := make(map[int]int)

	for _, val := range col {
		_, contains := numbers[val]

		if contains {
			return false
		}

		if val != 0 {
			numbers[val] = 1
		}
	}

	return true
}
