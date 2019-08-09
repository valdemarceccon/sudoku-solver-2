package solver

type Line [9]int
type Puzzle [9]Line

type Solver struct {
	puzzle Puzzle
}

func New(puzzle Puzzle) *Solver {
	return &Solver{puzzle: puzzle}
}

func (p *Solver) LineAt(at int) Line {
	return p.puzzle[at]
}

func (p *Solver) ColumnAt(at int) (column Line) {
	for idx, val := range p.puzzle {
		column[idx] = val[at]
	}
	return
}
func (p *Solver) QuadrantAt(startX, startY int) (quadrant Line) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			quadrant[x*3+y] = p.puzzle[x+startX][y+startY]
		}
	}
	return
}
