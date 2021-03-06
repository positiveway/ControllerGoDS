//go:build windows

package platformSpecific

import (
	"github.com/go-vgo/robotgo"
)

func CloseInputResources() {}

func InitInput() {
	robotgo.MouseSleep = 25
	robotgo.KeySleep = 50
}

func PressKeyOrMouse(key int) {
	switch key {
	case LeftMouse:
		robotgo.Toggle("left")
	case RightMouse:
		robotgo.Toggle("right")
	case MiddleMouse:
		robotgo.Toggle("center") //TODO: check
	default:
		robotgo.KeyToggle(string(key))
	}
}

//type int32 = int

func ReleaseKeyOrMouse(key int) {
	switch key {
	case LeftMouse:
		robotgo.Toggle("left", "up")
	case RightMouse:
		robotgo.Toggle("right", "up")
	case MiddleMouse:
		robotgo.Toggle("center", "up") //TODO: check
	default:
		robotgo.KeyToggle(string(key), "up")
	}
}

func TypeKey(key int) {
	robotgo.KeyTap(string(key))
}

func MoveMouse(x, y int32) {
	robotgo.MoveRelative(int(x), int(y))
}

func ScrollHorizontal(direction int32) {
	robotgo.Scroll(int(direction), 0)
}

func ScrollVertical(direction int32) {
	robotgo.Scroll(0, int(direction))
}
