package main

import (
	"fmt"
	"math"
	"os"
)

var (
	ColorBlack = "\033[30m"
	ColorWhite = "\033[37m"

	ColorRed     = "\033[31m" // 625 - 740
	ColorYellow  = "\033[33m" // 565 - 590
	ColorGreen   = "\033[32m" // 520 - 565
	ColorCyan    = "\033[36m" // 500 - 520
	ColorBlue    = "\033[34m" // 435 - 500
	ColorMagenta = "\033[35m" // 380 - 435

	ColorReset = "\033[0m"

	Clear = "\033[2J"

	Corner         = "+"
	HorizontalLine = "-"
	VerticalLine   = "|"
)

func DrawInit() {
	fmt.Println(Clear)
	DrawLegend()
}

func Draw(r *Relaxation) {
	// fmt.Println(Clear)
	DrawTitle()
	// DrawText()
	DrawGrid(r)
	Jump(72, 24)
	fmt.Println()
}

func DrawLegend() {
	Jump(4, 24)
	fmt.Print(ColorWhite, "Legend: ", ColorMagenta, "Boundary ", GetColor(0), "01", GetColor(2), "23", GetColor(4), "45", GetColor(6), "67", GetColor(8), "89")
}

func DrawTitle() {
	Jump(20, 1)
	fmt.Println(ColorWhite, "Method of Relaxation - 1/10/2021 - c0nrad", ColorReset)
}

func DrawGrid(r *Relaxation) {
	GridOffsetX := 4
	GridOffsetY := 3

	for x := 0; x < r.Width; x++ {
		for y := 0; y < r.Height; y++ {
			Jump(x+GridOffsetX, y+GridOffsetY)

			v := int(math.Round(r.At(x, y)))
			if r.IsBoundaryCondition(x, y) {
				fmt.Print(ColorMagenta, v, ColorReset)
			} else {

				fmt.Print(GetColor(v), v, ColorReset)
			}
		}
	}
}

func GetColor(v int) string {
	if v == 0 || v == 1 {
		return ColorWhite
	}

	if v == 9 || v == 8 {
		return ColorRed
	}

	if v == 7 || v == 6 {
		return ColorYellow
	}

	if v == 5 || v == 4 {
		return ColorYellow
	}

	if v == 3 || v == 2 {
		return ColorGreen
	}

	panic("not possible")
}

func Jump(x, y int) {
	os.Stdout.WriteString(fmt.Sprintf("\033[%d;%df", y, x))
}
