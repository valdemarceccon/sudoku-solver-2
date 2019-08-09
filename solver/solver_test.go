package solver_test

import (
	"reflect"
	"sudoku-solver/solver"
	"testing"
)

func TestConstraints(t *testing.T) {

	solvedPuzzle := solver.Puzzle{
		solver.Line{7, 3, 9, 2, 5, 4, 6, 8, 1},
		solver.Line{4, 6, 8, 3, 9, 1, 5, 7, 2},
		solver.Line{5, 2, 1, 8, 6, 7, 3, 4, 9},
		solver.Line{2, 1, 5, 7, 3, 6, 4, 9, 8},
		solver.Line{8, 9, 3, 5, 4, 2, 1, 6, 7},
		solver.Line{6, 4, 7, 9, 1, 8, 2, 3, 5},
		solver.Line{1, 8, 2, 6, 7, 3, 9, 5, 4},
		solver.Line{9, 7, 6, 4, 2, 5, 8, 1, 3},
		solver.Line{3, 5, 4, 1, 8, 9, 7, 2, 6},
	}

	t.Run("given a puzzle, get needed row", func(t *testing.T) {
		sol := solver.New(solvedPuzzle)

		gotLine := sol.LineAt(1)
		expectedLine := solver.Line{4, 6, 8, 3, 9, 1, 5, 7, 2}

		assertLine(t, gotLine, expectedLine)
	})

	t.Run("given a puzzle, get needed column", func(t *testing.T) {
		sol := solver.New(solvedPuzzle)

		gotLine := sol.ColumnAt(1)
		expectedLine := solver.Line{3, 6, 2, 1, 9, 4, 8, 7, 5}

		assertLine(t, gotLine, expectedLine)
	})

	t.Run("given a puzzle, get second 'quadrant'", func(t *testing.T) {
		sol := solver.New(solvedPuzzle)

		gotQuadrant := sol.QuadrantAt(0, 6)
		expectedQuadrant := solver.Line{6, 8, 1, 5, 7, 2, 3, 4, 9}

		assertLine(t, gotQuadrant, expectedQuadrant)
	})

	t.Run("given a puzzle, get middle 'quadrant'", func(t *testing.T) {
		sol := solver.New(solvedPuzzle)

		gotQuadrant := sol.QuadrantAt(3, 3)
		expectedQuadrant := solver.Line{7, 3, 6, 5, 4, 2, 9, 1, 8}

		assertLine(t, gotQuadrant, expectedQuadrant)
	})
}

func assertLine(t *testing.T, gotLine solver.Line, expectedLine solver.Line) {
	if !reflect.DeepEqual(gotLine, expectedLine) {
		t.Errorf("want %v got %v", expectedLine, gotLine)
	}
}
