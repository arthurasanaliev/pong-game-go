package game

import "image/color"

type Object struct {
	x, y          float32
	width, height float32
	color         color.Color
}
