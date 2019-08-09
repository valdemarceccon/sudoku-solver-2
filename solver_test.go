package sudoku_solver_test

import (
	"reflect"
	sudoku_solver "sudoku-solver"
	"testing"
)

func TestConstraints(t *testing.T) {

	solvedPuzzle := sudoku_solver.Puzzle{
		sudoku_solver.Line{7, 3, 9, 2, 5, 4, 6, 8, 1},
		sudoku_solver.Line{4, 6, 8, 3, 9, 1, 5, 7, 2},
		sudoku_solver.Line{5, 2, 1, 8, 6, 7, 3, 4, 9},
		sudoku_solver.Line{2, 1, 5, 7, 3, 6, 4, 9, 8},
		sudoku_solver.Line{8, 9, 3, 5, 4, 2, 1, 6, 7},
		sudoku_solver.Line{6, 4, 7, 9, 1, 8, 2, 3, 5},
		sudoku_solver.Line{1, 8, 2, 6, 7, 3, 9, 5, 4},
		sudoku_solver.Line{9, 7, 6, 4, 2, 5, 8, 1, 3},
		sudoku_solver.Line{3, 5, 4, 1, 8, 9, 7, 2, 6},
	}

	t.Run("given a puzzle, get needed row", func(t *testing.T) {
		solver := sudoku_solver.Solver{Puzzle: solvedPuzzle}

		gotLine := solver.Puzzle.LineAt(1)
		expectedLine := sudoku_solver.Line{4, 6, 8, 3, 9, 1, 5, 7, 2}

		if !reflect.DeepEqual(gotLine, expectedLine) {
			t.Errorf("want %v got %v", gotLine, expectedLine)
		}

	})

	t.Run("given a puzzle, get needed column", func(t *testing.T) {
		solver := sudoku_solver.Solver{Puzzle: solvedPuzzle}

		gotLine := solver.Puzzle.ColumnAt(1)

		expectedLine := sudoku_solver.Line{3, 6, 2, 1, 9, 4, 8, 7, 5}

		if !reflect.DeepEqual(gotLine, expectedLine) {
			t.Errorf("want %v got %v", gotLine, expectedLine)
		}

	})
}
