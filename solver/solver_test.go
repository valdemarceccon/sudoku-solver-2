package solver_test

import (
	"reflect"
	"sudoku-solver/solver"
	"testing"
)

func TestGroupGetters(t *testing.T) {
	var solvedPuzzle = solver.Sudoku{
		{7, 3, 9, 2, 5, 4, 6, 8, 1},
		{4, 6, 8, 3, 9, 1, 5, 7, 2},
		{5, 2, 1, 8, 6, 7, 3, 4, 9},
		{2, 1, 5, 7, 3, 6, 4, 9, 8},
		{8, 9, 3, 5, 4, 2, 1, 6, 7},
		{6, 4, 7, 9, 1, 8, 2, 3, 5},
		{1, 8, 2, 6, 7, 3, 9, 5, 4},
		{9, 7, 6, 4, 2, 5, 8, 1, 3},
		{3, 5, 4, 1, 8, 9, 7, 2, 6},
	}

	t.Run("given a puzzle, return first line", func(t *testing.T) {
		expectedRow := []int{7, 3, 9, 2, 5, 4, 6, 8, 1}

		rowGot := solvedPuzzle.RowAt(0)

		assertSliceEquals(t, expectedRow, rowGot)
	})

	t.Run("given a puzzle, return middle line", func(t *testing.T) {
		expectedRow := []int{2, 1, 5, 7, 3, 6, 4, 9, 8}

		rowGot := solvedPuzzle.RowAt(3)

		assertSliceEquals(t, expectedRow, rowGot)
	})

	t.Run("given a puzzle, return second column", func(t *testing.T) {
		expectedColumn := []int{3, 6, 2, 1, 9, 4, 8, 7, 5}

		columnGot := solvedPuzzle.ColumnAt(1)

		assertSliceEquals(t, expectedColumn, columnGot)
	})

	t.Run("given a puzzle, return fifth column", func(t *testing.T) {
		expectedColumn := []int{5, 9, 6, 3, 4, 1, 7, 2, 8}

		columnGot := solvedPuzzle.ColumnAt(4)

		assertSliceEquals(t, expectedColumn, columnGot)
	})

	t.Run("given a puzzle, return middle quadrant", func(t *testing.T) {
		expectedQuadrant := []int{7, 3, 6, 5, 4, 2, 9, 1, 8}

		quadrantGot := solvedPuzzle.QuadrantAt(3, 3)

		assertSliceEquals(t, expectedQuadrant, quadrantGot)
	})

	t.Run("given a puzzle, return second quadrant from the last row", func(t *testing.T) {
		expectedQuadrant := []int{6, 7, 3, 4, 2, 5, 1, 8, 9}

		quadrantGot := solvedPuzzle.QuadrantAt(7, 4)

		assertSliceEquals(t, expectedQuadrant, quadrantGot)
	})
}

func TestConstraintsChecker(t *testing.T) {
	var validAndInvalidGroups = solver.Sudoku{
		{7, 3, 9, 0, 0, 0, 0, 0, 0},
		{4, 6, 8, 0, 0, 8, 0, 0, 0},
		{5, 2, 1, 0, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 0, 0},
		{8, 9, 3, 5, 4, 2, 1, 6, 7},
		{0, 4, 7, 0, 0, 0, 0, 0, 0},
		{1, 8, 2, 6, 7, 3, 9, 5, 4},
		{0, 7, 6, 0, 0, 0, 0, 0, 0},
		{0, 5, 8, 0, 0, 0, 0, 0, 0},
	}
	t.Run("when there is a invalid column, checker return a error", func(t *testing.T) {
		wanted := solver.InvalidPuzzleError.Error()

		err := validAndInvalidGroups.CheckValidColumn(1)

		assertError(t, wanted, err)
	})

	t.Run("when there is a valid column but incomplete, checker return no error", func(t *testing.T) {
		err := validAndInvalidGroups.CheckValidColumn(0)

		assertNotError(t, err)
	})

	t.Run("when there is a valid row but incomplete, checker return no error", func(t *testing.T) {
		err := validAndInvalidGroups.CheckValidRow(0)

		assertNotError(t, err)
	})

	t.Run("when there is a invalid row, checker return no error", func(t *testing.T) {
		wanted := solver.InvalidPuzzleError.Error()

		err := validAndInvalidGroups.CheckValidRow(1)

		assertError(t, wanted, err)
	})

	t.Run("when there is a invalid quadrant, checker return a error", func(t *testing.T) {
		wanted := solver.InvalidPuzzleError.Error()

		err := validAndInvalidGroups.CheckValidQuadrant(7, 1)

		assertError(t, wanted, err)
	})

	t.Run("when there is a valid quadrant, checker return no error", func(t *testing.T) {
		err := validAndInvalidGroups.CheckValidQuadrant(2, 0)

		assertNotError(t, err)
	})
}

func TestPuzzleSolver(t *testing.T) {
	var almostSolved = solver.Sudoku{
		{7, 3, 9, 0, 5, 4, 6, 8, 1},
		{4, 6, 8, 3, 9, 1, 5, 7, 2},
		{5, 2, 1, 8, 0, 7, 3, 4, 9},
		{2, 1, 5, 7, 3, 6, 4, 9, 8},
		{8, 9, 3, 5, 4, 2, 1, 6, 7},
		{6, 4, 7, 9, 1, 0, 2, 3, 5},
		{1, 8, 2, 6, 7, 3, 9, 5, 4},
		{9, 7, 6, 4, 2, 5, 8, 1, 3},
		{3, 5, 4, 1, 8, 9, 7, 2, 6},
	}
	var solved = solver.Sudoku{
		{7, 3, 9, 2, 5, 4, 6, 8, 1},
		{4, 6, 8, 3, 9, 1, 5, 7, 2},
		{5, 2, 1, 8, 6, 7, 3, 4, 9},
		{2, 1, 5, 7, 3, 6, 4, 9, 8},
		{8, 9, 3, 5, 4, 2, 1, 6, 7},
		{6, 4, 7, 9, 1, 8, 2, 3, 5},
		{1, 8, 2, 6, 7, 3, 9, 5, 4},
		{9, 7, 6, 4, 2, 5, 8, 1, 3},
		{3, 5, 4, 1, 8, 9, 7, 2, 6},
	}
	var impossible = solver.Sudoku{
		{7, 3, 9, 0, 5, 4, 6, 8, 1},
		{4, 6, 8, 2, 9, 1, 5, 7, 2},
		{5, 2, 1, 8, 0, 7, 3, 4, 9},
		{2, 1, 5, 7, 3, 6, 4, 9, 8},
		{8, 9, 3, 5, 4, 2, 1, 6, 7},
		{6, 4, 7, 9, 1, 0, 2, 3, 5},
		{1, 8, 2, 6, 7, 3, 9, 5, 4},
		{9, 7, 6, 4, 2, 5, 8, 1, 3},
		{3, 5, 4, 1, 8, 9, 7, 2, 6},
	}

	t.Run("given a incomplete puzzle, return next empty position and possibilities", func(t *testing.T) {
		toBeSolved := almostSolved.FindAllEmpty()

		gotLen := len(toBeSolved)

		testData := []struct {
			x, y   int
			values []int
		}{
			{x: 0, y: 3, values: []int{2}},
			{x: 2, y: 4, values: []int{6}},
			{x: 5, y: 5, values: []int{8}},
		}

		wantedLen := len(testData)
		if gotLen != wantedLen {
			t.Errorf("should have encountered %d, but found only %d", gotLen, wantedLen)
		}

		for k, v := range testData {
			if toBeSolved[k].X != v.x || toBeSolved[k].Y != v.y || !reflect.DeepEqual(toBeSolved[k].Possibles, v.values) {
				t.Errorf("found data diverge from expected, wanted %v, got %v", v, toBeSolved[k])
			}
		}
	})

	t.Run("given a impossible puzzle, returns a possible with empty possibilities", func(t *testing.T) {
		allEmpty := impossible.FindAllEmpty()
		wantLen := 3

		gotLen := len(allEmpty)
		if gotLen != wantLen {
			t.Errorf("want len %d, but got %d", wantLen, gotLen)
		}

		assertThatHaveEmptyPossibility(t, allEmpty)
	})

	t.Run("given a solved puzzle, find all empty should return a empty array", func(t *testing.T) {
		empty := solved.FindAllEmpty()
		wantLen := 0
		gotLen := len(empty)

		if wantLen != gotLen {
			t.Errorf("wanted %d, but got %d", gotLen, wantLen)
		}
	})

	t.Run("given a not solved, but possible puzzle find de complete puzzle", func(t *testing.T) {
		testAlmostSolved := almostSolved.Copy()
		testAlmostSolved.Solve()

		if !reflect.DeepEqual(solved, testAlmostSolved) {
			t.Errorf("should have resolved the puzzle. gotten result %v", testAlmostSolved)
		}
	})

}

func assertThatHaveEmptyPossibility(t *testing.T, pendents []solver.Pendent) {
	t.Helper()
	for _, v := range pendents {
		if len(v.Possibles) == 0 {
			return
		}
	}
	t.Error("should have a empty possibility")
}

func assertNotError(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Errorf("was expecting no error, but got %v", e.Error())
	}
}

func assertSliceEquals(t *testing.T, wanted, got []int) {
	t.Helper()

	if !reflect.DeepEqual(wanted, got) {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}

func assertError(t *testing.T, expected string, errGot error) {
	t.Helper()

	if errGot == nil {
		t.Fatal("was expecting an error, but got none")
	}

	if errGot.Error() != expected {
		t.Errorf("was expecting %v, got %v", expected, errGot.Error())
	}
}
