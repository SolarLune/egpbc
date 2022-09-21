package egpbc

import "github.com/hajimehoshi/ebiten/v2"

type microsoftConstantDefinition struct {
	A, B, X, Y ebiten.StandardGamepadButton
	Menu       ebiten.StandardGamepadButton
	View       ebiten.StandardGamepadButton
	Xbox       ebiten.StandardGamepadButton
	DPadUp     ebiten.StandardGamepadButton
	DPadRight  ebiten.StandardGamepadButton
	DPadDown   ebiten.StandardGamepadButton
	DPadLeft   ebiten.StandardGamepadButton

	LeftBumper   ebiten.StandardGamepadButton
	RightBumper  ebiten.StandardGamepadButton
	LeftTrigger  ebiten.StandardGamepadButton
	RightTrigger ebiten.StandardGamepadButton
	LeftStick    ebiten.StandardGamepadButton
	RightStick   ebiten.StandardGamepadButton

	LeftStickX  ebiten.StandardGamepadAxis
	LeftStickY  ebiten.StandardGamepadAxis
	RightStickX ebiten.StandardGamepadAxis
	RightStickY ebiten.StandardGamepadAxis

	StickLeft  int
	StickRight int
	StickUp    int
	StickDown  int
}

func (cs microsoftConstantDefinition) ButtonCodeFromString(buttonCode string) ebiten.StandardGamepadButton {

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
	case "Menu":
		return cs.Menu
	case "Select":
		fallthrough
	case "View":
		return cs.View
	case "Home":
		fallthrough
	case "Xbox":
		return cs.Xbox
	case "D-Pad Up":
		return cs.DPadUp
	case "D-Pad Right":
		return cs.DPadRight
	case "D-Pad Down":
		return cs.DPadDown
	case "D-Pad Left":
		return cs.DPadLeft
	case "Left Bumper":
		return cs.LeftBumper
	case "Left Stick":
		return cs.LeftStick
	case "Left Trigger":
		return cs.LeftTrigger
	case "Right Bumper":
		return cs.RightBumper
	case "Right Stick":
		return cs.RightStick
	case "Right Trigger":
		return cs.RightTrigger
	}

	return -1

}

func (cs microsoftConstantDefinition) StringFromButtonCode(buttonCode ebiten.StandardGamepadButton) string {

	switch buttonCode {

	case cs.A:
		return "A"
	case cs.B:
		return "B"
	case cs.X:
		return "X"
	case cs.Y:
		return "Y"
	case cs.Menu:
		return "Menu"
	case cs.View:
		return "View"
	case cs.Xbox:
		return "Xbox"
	case cs.DPadUp:
		return "D-Pad Up"
	case cs.DPadRight:
		return "D-Pad Right"
	case cs.DPadDown:
		return "D-Pad Down"
	case cs.DPadLeft:
		return "D-Pad Left"
	case cs.LeftBumper:
		return "Left Bumper"
	case cs.LeftStick:
		return "Left Stick"
	case cs.LeftTrigger:
		return "Left Trigger"
	case cs.RightBumper:
		return "Right Bumper"
	case cs.RightStick:
		return "Right Stick"
	case cs.RightTrigger:
		return "Right Trigger"
	}

	return ""

}

func (cs microsoftConstantDefinition) AxisCodeFromString(axisCode string) ebiten.StandardGamepadAxis {

	switch axisCode {

	case "Left Stick X":
		return cs.LeftStickX
	case "Left Stick Y":
		return cs.LeftStickY
	case "Right Stick X":
		return cs.RightStickX
	case "Right Stick Y":
		return cs.RightStickY
	}

	return -1

}

func (cs microsoftConstantDefinition) StringFromAxisCode(axisCode ebiten.StandardGamepadAxis) string {

	switch axisCode {
	case cs.LeftStickX:
		return "Left Stick X"
	case cs.LeftStickY:
		return "Left Stick Y"
	case cs.RightStickX:
		return "Right Stick X"
	case cs.RightStickY:
		return "Right Stick Y"
	}

	return ""

}

var MicrosoftConstants = microsoftConstantDefinition{
	A:            ebiten.StandardGamepadButtonRightBottom,
	B:            ebiten.StandardGamepadButtonRightRight,
	X:            ebiten.StandardGamepadButtonRightLeft,
	Y:            ebiten.StandardGamepadButtonRightTop,
	Menu:         ebiten.StandardGamepadButtonCenterRight,
	View:         ebiten.StandardGamepadButtonCenterLeft,
	Xbox:         ebiten.StandardGamepadButtonCenterCenter,
	DPadUp:       ebiten.StandardGamepadButtonLeftTop,
	DPadRight:    ebiten.StandardGamepadButtonLeftRight,
	DPadDown:     ebiten.StandardGamepadButtonLeftBottom,
	DPadLeft:     ebiten.StandardGamepadButtonLeftLeft,
	LeftStick:    ebiten.StandardGamepadButtonLeftStick,
	RightStick:   ebiten.StandardGamepadButtonRightStick,
	LeftBumper:   ebiten.StandardGamepadButtonFrontTopLeft,
	RightBumper:  ebiten.StandardGamepadButtonFrontTopRight,
	LeftTrigger:  ebiten.StandardGamepadButtonFrontBottomLeft,
	RightTrigger: ebiten.StandardGamepadButtonFrontBottomRight,

	LeftStickX:  ebiten.StandardGamepadAxisLeftStickHorizontal,
	LeftStickY:  ebiten.StandardGamepadAxisLeftStickVertical,
	RightStickX: ebiten.StandardGamepadAxisRightStickHorizontal,
	RightStickY: ebiten.StandardGamepadAxisRightStickVertical,

	StickLeft:  -1,
	StickRight: 1,
	StickUp:    -1,
	StickDown:  1,
}
