package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	Object
}

func NewPaddle(x, y, width, height float32, color color.Color) *Paddle {
	return &Paddle{
		Object: Object{
			x:      x,
			y:      y,
			width:  width,
			height: height,
			color:  color,
		},
	}
}

func (p *Paddle) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if p.y-PADDLE_SPEED >= 0 {
			p.y -= PADDLE_SPEED
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if p.y+PADDLE_SPEED+p.height < SCREEN_HEIGHT {
			p.y += PADDLE_SPEED
		}
	}
}
