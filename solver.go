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
func (p Puzzle) QuadrantAt(startX, startY int) (quadrant Line) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			quadrant[x*3+y] = p[x+startX][y+startY]
		}
	}
	return
}
