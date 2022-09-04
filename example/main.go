package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/egpbc"
)

type Game struct{}

func (game *Game) Update() error {

	// easy shortcutting
	xbox := egpbc.MicrosoftConstants

	for i := 0; i < 16; i++ {

		buttonID := ebiten.StandardGamepadButton(i)

		if ebiten.IsStandardGamepadButtonAvailable(0, buttonID) && ebiten.IsStandardGamepadButtonPressed(0, buttonID) {
			fmt.Println("Pressed: ", xbox.StringFromButtonCode(buttonID))
		}

	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {}

func (game *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}

}
