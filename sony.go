package egpbc

import "github.com/hajimehoshi/ebiten/v2"

type sonyConstantDefinition struct {
	Cross     ebiten.StandardGamepadButton
	Circle    ebiten.StandardGamepadButton
	Square    ebiten.StandardGamepadButton
	Triangle  ebiten.StandardGamepadButton
	Options   ebiten.StandardGamepadButton
	Create    ebiten.StandardGamepadButton
	PS        ebiten.StandardGamepadButton
	DPadUp    ebiten.StandardGamepadButton
	DPadRight ebiten.StandardGamepadButton
	DPadDown  ebiten.StandardGamepadButton
	DPadLeft  ebiten.StandardGamepadButton

	L1 ebiten.StandardGamepadButton
	R1 ebiten.StandardGamepadButton
	L2 ebiten.StandardGamepadButton
	R2 ebiten.StandardGamepadButton
	L3 ebiten.StandardGamepadButton
	R3 ebiten.StandardGamepadButton

	LeftStickX  ebiten.StandardGamepadAxis
	LeftStickY  ebiten.StandardGamepadAxis
	RightStickX ebiten.StandardGamepadAxis
	RightStickY ebiten.StandardGamepadAxis

	StickLeft  int
	StickRight int
	StickUp    int
	StickDown  int
}

func (cs sonyConstantDefinition) ButtonCodeFromString(buttonCode string) ebiten.StandardGamepadButton {

	switch buttonCode {

	case "Cross":
		return cs.Cross
	case "Circle":
		return cs.Circle
	case "Square":
		return cs.Square
	case "Triangle":
		return cs.Triangle
	case "Start":
		fallthrough
	case "Options":
		return cs.Options
	case "Select": // TODO: Confirm old-school Select button is the new Create button?
		fallthrough
	case "Create":
		return cs.Create
	case "Home":
		fallthrough
	case "PS":
		return cs.PS
	case "D-Pad Up":
		return cs.DPadUp
	case "D-Pad Right":
		return cs.DPadRight
	case "D-Pad Down":
		return cs.DPadDown
	case "D-Pad Left":
		return cs.DPadLeft
	case "L1":
		return cs.L1
	case "L2":
		return cs.L2
	case "L3":
		return cs.L3
	case "R1":
		return cs.R1
	case "R2":
		return cs.R2
	case "R3":
		return cs.R3
	}

	return -1

}

func (cs sonyConstantDefinition) StringFromButtonCode(buttonCode ebiten.StandardGamepadButton) string {

	switch buttonCode {

	case cs.Cross:
		return "Cross"
	case cs.Circle:
		return "Circle"
	case cs.Triangle:
		return "Triangle"
	case cs.Square:
		return "Square"
	case cs.Options:
		return "Options"
	case cs.Create:
		return "Create"
	case cs.PS:
		return "PS"
	case cs.DPadUp:
		return "D-Pad Up"
	case cs.DPadRight:
		return "D-Pad Right"
	case cs.DPadDown:
		return "D-Pad Down"
	case cs.DPadLeft:
		return "D-Pad Left"
	case cs.L1:
		return "L1"
	case cs.L2:
		return "L2"
	case cs.L3:
		return "L3"
	case cs.R1:
		return "R1"
	case cs.R2:
		return "R2"
	case cs.R3:
		return "R3"
	}

	return ""

}

func (cs sonyConstantDefinition) AxisCodeFromString(axisCode string) ebiten.StandardGamepadAxis {

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

func (cs sonyConstantDefinition) StringFromAxisCode(axisCode ebiten.StandardGamepadAxis) string {

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

var SonyConstants = sonyConstantDefinition{
	Cross:     ebiten.StandardGamepadButtonRightBottom,
	Circle:    ebiten.StandardGamepadButtonRightRight,
	Square:    ebiten.StandardGamepadButtonRightLeft,
	Triangle:  ebiten.StandardGamepadButtonRightTop,
	Options:   ebiten.StandardGamepadButtonCenterRight,
	Create:    ebiten.StandardGamepadButtonCenterLeft,
	PS:        ebiten.StandardGamepadButtonCenterCenter,
	DPadUp:    ebiten.StandardGamepadButtonLeftTop,
	DPadRight: ebiten.StandardGamepadButtonLeftRight,
	DPadDown:  ebiten.StandardGamepadButtonLeftBottom,
	DPadLeft:  ebiten.StandardGamepadButtonLeftLeft,
	L3:        ebiten.StandardGamepadButtonLeftStick,
	R3:        ebiten.StandardGamepadButtonRightStick,
	L1:        ebiten.StandardGamepadButtonFrontTopLeft,
	R1:        ebiten.StandardGamepadButtonFrontTopRight,
	L2:        ebiten.StandardGamepadButtonFrontBottomLeft,
	R2:        ebiten.StandardGamepadButtonFrontBottomRight,

	LeftStickX:  ebiten.StandardGamepadAxisLeftStickHorizontal,
	LeftStickY:  ebiten.StandardGamepadAxisLeftStickVertical,
	RightStickX: ebiten.StandardGamepadAxisRightStickHorizontal,
	RightStickY: ebiten.StandardGamepadAxisRightStickVertical,

	StickLeft:  -1,
	StickRight: 1,
	StickUp:    -1,
	StickDown:  1,
}
