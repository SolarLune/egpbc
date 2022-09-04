package egpbc

import "github.com/hajimehoshi/ebiten/v2"

type nintendoConstantDefinition struct {
	A, B, X, Y ebiten.StandardGamepadButton
	Plus       ebiten.StandardGamepadButton
	Minus      ebiten.StandardGamepadButton
	Home       ebiten.StandardGamepadButton
	DPadUp     ebiten.StandardGamepadButton
	DPadRight  ebiten.StandardGamepadButton
	DPadDown   ebiten.StandardGamepadButton
	DPadLeft   ebiten.StandardGamepadButton

	L          ebiten.StandardGamepadButton
	R          ebiten.StandardGamepadButton
	ZL         ebiten.StandardGamepadButton
	ZR         ebiten.StandardGamepadButton
	LeftStick  ebiten.StandardGamepadButton
	RightStick ebiten.StandardGamepadButton

	LeftStickX  ebiten.StandardGamepadAxis
	LeftStickY  ebiten.StandardGamepadAxis
	RightStickX ebiten.StandardGamepadAxis
	RightStickY ebiten.StandardGamepadAxis

	StickLeft  int
	StickRight int
	StickUp    int
	StickDown  int
}

func (cs nintendoConstantDefinition) ButtonCodeFromString(buttonCode string) ebiten.StandardGamepadButton {

	switch buttonCode {

	case "A":
		return cs.A
	case "B":
		return cs.B
	case "X":
		return cs.X
	case "Y":
		return cs.Y
	case "Start":
		fallthrough
	case "Plus":
		return cs.Plus
	case "Select":
		fallthrough
	case "Minus":
		return cs.Minus
	case "Home":
		return cs.Home
	case "D-Pad Up":
		return cs.DPadUp
	case "D-Pad Right":
		return cs.DPadRight
	case "D-Pad Down":
		return cs.DPadDown
	case "D-Pad Left":
		return cs.DPadLeft
	case "L":
		return cs.L
	case "Left Stick":
		return cs.LeftStick
	case "ZL":
		return cs.ZL
	case "R":
		return cs.R
	case "Right Stick":
		return cs.RightStick
	case "ZR":
		return cs.ZR
	}

	return -1

}

func (cs nintendoConstantDefinition) StringFromButtonCode(buttonCode ebiten.StandardGamepadButton) string {

	switch buttonCode {

	case cs.A:
		return "A"
	case cs.B:
		return "B"
	case cs.X:
		return "X"
	case cs.Y:
		return "Y"
	case cs.Plus:
		return "Plus"
	case cs.Minus:
		return "Minus"
	case cs.Home:
		return "Home"
	case cs.DPadUp:
		return "D-Pad Up"
	case cs.DPadRight:
		return "D-Pad Right"
	case cs.DPadDown:
		return "D-Pad Down"
	case cs.DPadLeft:
		return "D-Pad Left"
	case cs.L:
		return "L"
	case cs.LeftStick:
		return "Left Stick"
	case cs.ZL:
		return "ZL"
	case cs.R:
		return "R"
	case cs.RightStick:
		return "Right Stick"
	case cs.ZR:
		return "ZR"
	}

	return ""

}

func (cs nintendoConstantDefinition) AxisCodeFromString(axisCode string) ebiten.StandardGamepadAxis {

	switch axisCode {

	case "left stick x":
		return cs.LeftStickX
	case "left stick y":
		return cs.LeftStickY
	case "right stick x":
		return cs.RightStickX
	case "right stick y":
		return cs.RightStickY
	}

	return -1

}

var NintendoConstants = nintendoConstantDefinition{
	A:          ebiten.StandardGamepadButtonRightRight,
	B:          ebiten.StandardGamepadButtonRightBottom,
	X:          ebiten.StandardGamepadButtonRightTop,
	Y:          ebiten.StandardGamepadButtonRightLeft,
	Plus:       ebiten.StandardGamepadButtonCenterRight,
	Minus:      ebiten.StandardGamepadButtonCenterLeft,
	Home:       ebiten.StandardGamepadButtonCenterCenter,
	DPadUp:     ebiten.StandardGamepadButtonLeftTop,
	DPadRight:  ebiten.StandardGamepadButtonLeftRight,
	DPadDown:   ebiten.StandardGamepadButtonLeftBottom,
	DPadLeft:   ebiten.StandardGamepadButtonLeftLeft,
	LeftStick:  ebiten.StandardGamepadButtonLeftStick,
	RightStick: ebiten.StandardGamepadButtonRightStick,
	L:          ebiten.StandardGamepadButtonFrontTopLeft,
	R:          ebiten.StandardGamepadButtonFrontTopRight,
	ZL:         ebiten.StandardGamepadButtonFrontBottomLeft,
	ZR:         ebiten.StandardGamepadButtonFrontBottomRight,

	LeftStickX:  ebiten.StandardGamepadAxisLeftStickHorizontal,
	LeftStickY:  ebiten.StandardGamepadAxisLeftStickVertical,
	RightStickX: ebiten.StandardGamepadAxisRightStickHorizontal,
	RightStickY: ebiten.StandardGamepadAxisRightStickVertical,

	StickLeft:  -1,
	StickRight: 1,
	StickUp:    -1,
	StickDown:  1,
}
