package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	tileSize     = 48
	boardOffsetY = 80
	screenWidth  = boardSize * tileSize
	screenHeight = boardOffsetY + boardSize*tileSize
)

type Game struct {
	Match *Match
}

func NewGame() *Game {
	return &Game{
		Match: NewMatch(),
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.Match.EndTurn()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 22, 28, 255})

	// Draw board.
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			tileColor := color.RGBA{65, 70, 80, 255}
			if (x+y)%2 == 0 {
				tileColor = color.RGBA{75, 80, 90, 255}
			}
			vector.FillRect(
				screen,
				float32(x*tileSize),
				float32(boardOffsetY+y*tileSize),
				float32(tileSize-1),
				float32(tileSize-1),
				tileColor,
				false,
			)
		}
	}

	// Draw heroes
	heroColors := [2]color.RGBA{
		{78, 158, 255, 255},
		{240, 90, 90, 255},
	}

	for i, hero := range g.Match.Heroes {
		x := float32(hero.X * tileSize)
		y := float32(boardOffsetY + hero.Y*tileSize)

		// white border identifies the active hero.
		if i == g.Match.Active {
			vector.FillRect(
				screen,
				x+5, y+5,
				float32(tileSize-10),
				float32(tileSize-10),
				color.White,
				false,
			)
		}

		vector.FillRect(
			screen, x+9, y+9,
			float32(tileSize-18),
			float32(tileSize-18),
			heroColors[i],
			false,
		)
	}

	hero := g.Match.Heroes[g.Match.Active]

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"Player %d's turn\nHP: %d | AP: %d | MP: %d\nSPACE: End turn",
		g.Match.Active+1,
		hero.HP,
		hero.AP,
		hero.MP,
	))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
