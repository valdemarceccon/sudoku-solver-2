package sudoku_solver

type Line [9]int
type Puzzle [9]Line

type Solver struct {
	Puzzle Puzzle
}

func (p Puzzle) LineAt(at int) Line {
	return p[at]
}

func (p Puzzle) ColumnAt(at int) (column Line) {
	for idx, val := range p {
		column[idx] = val[at]
	}
	return
}
func (p Puzzle) QuadrantAt(at int) (quadrant Line) {

	quadrant[0] = p[0][6]
	quadrant[1] = p[0][7]
	quadrant[2] = p[0][8]
	quadrant[3] = p[1][6]
	quadrant[4] = p[1][7]
	quadrant[5] = p[1][8]
	quadrant[6] = p[2][6]
	quadrant[7] = p[2][7]
	quadrant[8] = p[2][8]
	return
}
