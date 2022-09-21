package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/solarlune/egpbc"
	"golang.org/x/image/font/basicfont"
)

type Game struct{}

func (game *Game) Update() error {

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	txt := "Press an input on the gamepad:\n"

	// easy shortcutting
	xbox := egpbc.MicrosoftConstants

	for i := 0; i < 16; i++ {

		buttonID := ebiten.StandardGamepadButton(i)

		if ebiten.IsStandardGamepadButtonAvailable(0, buttonID) && ebiten.IsStandardGamepadButtonPressed(0, buttonID) {
			txt += "Pressed: " + xbox.StringFromButtonCode(buttonID) + "\n"
		}

	}

	lsX := ebiten.StandardGamepadAxisValue(0, xbox.LeftStickX)
	lsY := ebiten.StandardGamepadAxisValue(0, xbox.LeftStickY)

	if math.Abs(lsX) > 0.1 || math.Abs(lsY) > 0.1 {
		txt += fmt.Sprintf("Left Stick Tilted: %.2f, %.2f\n", lsX, lsY)
	}

	rsX := ebiten.StandardGamepadAxisValue(0, xbox.RightStickX)
	rsY := ebiten.StandardGamepadAxisValue(0, xbox.RightStickY)

	if math.Abs(rsX) > 0.1 || math.Abs(rsY) > 0.1 {
		txt += fmt.Sprintf("Right Stick Tilted: %.2f, %.2f\n", rsX, rsY)
	}

	text.Draw(screen, txt, basicfont.Face7x13, 20, 20, color.White)
}

func (game *Game) Layout(w, h int) (int, int) {
	return 320, 240
}

func main() {

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}

}
