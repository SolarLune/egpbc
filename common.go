package egpbc

func AxisIsLeft(axisValue float64) bool {
	return axisValue < 0
}

func AxisIsRight(axisValue float64) bool {
	return axisValue > 0
}

func AxisIsUp(axisValue float64) bool {
	return axisValue < 0
}

func AxisIsDown(axisValue float64) bool {
	return axisValue > 0
}
