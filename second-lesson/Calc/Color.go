package calc

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/fatih/structs"
)

type ColorStruct struct {
	White, Red, Yellow, Blue, Cyan, Green, Magneta, Black string
}

func RefStruct(c ColorStruct) []interface{} {
	c = ColorStruct{"white", "red", "yellow", "blue", "cyan", "green", "magneta", "black"}
	val := structs.Values(c)
	return val
}

func ChangeColor(x int) {

	col := RefStruct(ColorStruct{})
	if x > len(col) {
		x = x / 2
	}
	switch col[x] {
	case "white":
		color.Set(color.FgWhite)
	case "red":
		color.Set(color.FgRed)
	case "yellow":
		color.Set(color.FgYellow)
	case "blue":
		color.Set(color.FgBlue)
	case "cyan":
		color.Set(color.FgCyan)
	case "green":
		color.Set(color.FgGreen)
	case "magneta":
		color.Set(color.FgMagenta)
	case "black":
		color.Set(color.FgBlack)
	default:
		fmt.Println("We dont this Color in sturct")
	}
}
