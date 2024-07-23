package game

import (
	"fmt"
	"image/color"

	"golang.org/x/image/font/basicfont"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	paddle   *Paddle
	ball     *Ball
	score    int
	gameOver bool
}

func NewGame() *Game {
	paddle := NewPaddle(740, 360, 20, 70, color.White)
	ball := NewBall(0, 0, 15, 15, 1, 1, 5, color.White)
	return &Game{
		paddle:   paddle,
		ball:     ball,
		score:    0,
		gameOver: false,
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if g.StartOver() {
			g.paddle = NewPaddle(740, 360, 20, 70, color.White)
			g.ball = NewBall(0, 0, 15, 15, 1, 1, 5, color.White)
			g.score = 0
			g.gameOver = false
		}
		return nil
	}
	g.paddle.Move()
	g.ball.Move(g)
	g.score += 1
	g.ball.speed += 0.005
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		text.Draw(screen, "Game Over!", basicfont.Face7x13, 370, 270, color.White)
		text.Draw(screen, "Press Enter to Start Over!", basicfont.Face7x13, 315, 300, color.White)
		return
	}
	vector.DrawFilledRect(screen, g.paddle.x, g.paddle.y, g.paddle.width, g.paddle.height, g.paddle.color, false)
	vector.DrawFilledRect(screen, g.ball.x, g.ball.y, g.ball.width, g.ball.height, g.ball.color, false)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), basicfont.Face7x13, 20, 20, color.White)
}

func (g *Game) Layout(w, h int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) StartOver() bool {
	return ebiten.IsKeyPressed(ebiten.KeyEnter)
}
