package main

import (
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// Read imahe ftom file that already exists
	snakeHeadFile, err := os.Open("snake.png")
	if err != nil {
		panic(err)
	}
	defer snakeHeadFile.Close()

	snakeHeadImage, err := png.Decode(snakeHeadFile)
	if err != nil {
		panic(err)
	}

	snakeHead := ebiten.NewImageFromImage(snakeHeadImage)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("SnaKe")
	game := &Game{snakeHeadSprite: snakeHead}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	
}

type Game struct {
	snakeHeadSprite *ebiten.Image
}

func (g *Game) Update() error {
	// Update the logival state
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate((320/2)-20, (220/2)-20)
	screen.DrawImage(g.snakeHeadSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Return the fame logical screen size.
	// The screen is automatically scaled.
	return 320, 240
}
