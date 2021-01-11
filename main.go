package main

import (
	"fmt"
	"math"
)

var Grid [][]int

type Relaxation struct {
	Grid              [][]float64
	BoundaryCondition [][]bool

	Width, Height int
}

func NewRelaxation(width, height int) Relaxation {
	out := Relaxation{Width: width, Height: height}
	for w := 0; w < width; w++ {
		out.Grid = append(out.Grid, make([]float64, height))
		out.BoundaryCondition = append(out.BoundaryCondition, make([]bool, height))
	}
	return out
}

func (r Relaxation) At(x, y int) float64 {
	r.CheckBounds(x, y)
	return r.Grid[x][y]
}

func (r Relaxation) IsBoundaryCondition(x, y int) bool {
	r.CheckBounds(x, y)
	return r.BoundaryCondition[x][y]
}

func (r Relaxation) Set(x, y int, v float64) {
	r.CheckBounds(x, y)
	r.Grid[x][y] = v
}

func (r Relaxation) SetBoundaryCondition(x, y int, v float64) {
	r.Set(x, y, v)
	r.BoundaryCondition[x][y] = true
}

func (r Relaxation) InBounds(x, y int) bool {
	return x >= 0 && y >= 0 && y < r.Height && x < r.Width
}

func (r Relaxation) CheckBounds(x, y int) {
	if r.InBounds(x, y) {
		return
	}

	fmt.Println("Out of bounds", x, y, r.Width, r.Height)
	panic("abc")
}

func (r *Relaxation) RelaxStep(target float64) bool {
	isDone := true
	for x := 0; x < r.Width; x++ {
		for y := 0; y < r.Height; y++ {
			if r.IsBoundaryCondition(x, y) {
				continue
			}

			if r.InBounds(x-1, y) && r.InBounds(x+1, y) && r.InBounds(x, y-1) && r.InBounds(x, y+1) {
				v := r.At(x-1, y) + r.At(x+1, y) + r.At(x, y-1) + r.At(x, y+1)
				v /= 4

				if math.Abs(r.At(x, y)-v) > target {
					isDone = false
				}

				r.Set(x, y, v)
			}
		}
	}
	return isDone
}

func Demo1Relaxation() Relaxation {
	Width := 73
	Height := 20

	r := NewRelaxation(Width, Height)

	for w := 0; w < Width; w++ {
		r.SetBoundaryCondition(w, 0, 9)
		r.SetBoundaryCondition(w, Height-1, 0)
	}
	for h := 1; h < Height-1; h++ {
		r.SetBoundaryCondition(0, h, 0)
		r.SetBoundaryCondition(Width-1, h, 0)
	}
	return r
}

func main() {
	r := Demo1Relaxation()
	DrawInit()

	isDone := false
	for !isDone {
		Draw(&r)
		isDone = r.RelaxStep(.000001)
	}
}
