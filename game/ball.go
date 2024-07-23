package game

import (
	"image/color"
)

type Ball struct {
	Object
	directionX, directionY float32
	speed                  float32
}

func NewBall(x, y, width, height, directionX, directionY, speed float32, color color.Color) *Ball {
	return &Ball{
		Object: Object{
			x:      x,
			y:      y,
			width:  width,
			height: height,
			color:  color,
		},
		directionX: directionX,
		directionY: directionY,
		speed:      speed,
	}
}

func (b *Ball) Move(g *Game) {
	if b.x+b.directionX*b.speed >= g.paddle.x {
		if b.y >= g.paddle.y && b.y+b.height <= g.paddle.y+g.paddle.height {
			b.directionX *= -1
		}
	}
	if !inBoundsX(b.x + b.directionX*b.speed) {
		b.directionX *= -1
	} else if !inBoundsX(b.x + b.directionX*b.speed + b.width) {
		g.gameOver = true
		return
	}
	if !inBoundsY(b.y + b.directionY*b.speed) {
		b.directionY *= -1
	} else if !inBoundsY(b.y + b.directionY*b.speed + b.height) {
		b.directionY *= -1
	}
	b.x += b.directionX * b.speed
	b.y += b.directionY * b.speed
}

func inBoundsX(x float32) bool {
	return x >= 0 && x < SCREEN_WIDTH
}

func inBoundsY(y float32) bool {
	return y >= 0 && y < SCREEN_HEIGHT
}
