package mainLogic

import (
	"strconv"
	"strings"
)

type Event struct {
	eventType string
	btnOrAxis string
	value     float64
}

func (event *Event) update(msg string) {
	event.eventType, found = EventTypeMap[msg[0]]
	if !found {
		PanicMisspelled(string(msg[0]))
	}
	if event.eventType != EvConnected && event.eventType != EvDisconnected && event.eventType != EvDropped {
		event.btnOrAxis, found = BtnAxisMap[msg[1]]
		if !found {
			PanicMisspelled(string(msg[1]))
		}
		if event.eventType == EvAxisChanged || event.eventType == EvButtonChanged {
			event.value, err = strconv.ParseFloat(msg[2:], 32)
			CheckErr(err)
		}
	}
}

func (event *Event) print() {
	print("%s %s %0.2f",
		strings.TrimPrefix(event.eventType, "Ev"),
		strings.TrimPrefix(strings.TrimPrefix(event.btnOrAxis, "Btn"), "Axis"),
		event.value)
}

const (
	AxisLeftStickX  string = "AxisLeftStickX"
	AxisLeftStickY         = "AxisLeftStickY"
	AxisLeftZ              = "AxisLeftZ"
	AxisRightStickX        = "AxisRightStickX"
	AxisRightStickY        = "AxisRightStickY"
	AxisRightZ             = "AxisRightZ"
	AxisDPadX              = "AxisDPadX"
	AxisDPadY              = "AxisDPadY"
	AxisUnknown            = "AxisUnknown"
)

var _AxisMap = map[uint8]string{
	'u': AxisLeftStickX,
	'v': AxisLeftStickY,
	'w': AxisLeftZ,
	'x': AxisRightStickX,
	'y': AxisRightStickY,
	'z': AxisRightZ,
	'0': AxisDPadX,
	'1': AxisDPadY,
	'2': AxisUnknown,
}

const HoldSuffix = "Hold"

func addHoldSuffix(btn string) string {
	return btn + HoldSuffix
}

func removeHoldSuffix(btn string) string {
	return strings.TrimSuffix(btn, HoldSuffix)
}

const (
	BtnSouth         string = "South"
	BtnEast                 = "East"
	BtnNorth                = "North"
	BtnWest                 = "West"
	BtnC                    = "BtnC"
	BtnZ                    = "BtnZ"
	BtnLeftTrigger          = "LB"
	BtnLeftTrigger2         = "LT"
	BtnRightTrigger         = "RB"
	BtnRightTrigger2        = "RT"
	BtnSelect               = "Select"
	BtnStart                = "Start"
	BtnMode                 = "Mode"
	BtnLeftThumb            = "LeftThumb"
	BtnRightThumb           = "RightThumb"
	BtnDPadUp               = "DPadUp"
	BtnDPadDown             = "DPadDown"
	BtnDPadLeft             = "DPadLeft"
	BtnDPadRight            = "DPadRight"
	BtnUnknown              = "BtnUnknown"
)

type Synonyms = map[string]string

func genBtnSynonyms() Synonyms {
	synonyms := Synonyms{
		"LeftTrigger":   BtnLeftTrigger,
		"LeftTrigger2":  BtnLeftTrigger2,
		"RightTrigger":  BtnRightTrigger,
		"RightTrigger2": BtnRightTrigger2,
		"LeftStick":     BtnLeftThumb,
		"RightStick":    BtnRightThumb,
	}
	for key, val := range synonyms {
		synonyms[addHoldSuffix(key)] = addHoldSuffix(val)
	}
	return synonyms
}

var BtnSynonyms = genBtnSynonyms()

var AllOriginalButtons = []string{
	BtnSouth,
	BtnEast,
	BtnNorth,
	BtnWest,
	BtnC,
	BtnZ,
	BtnLeftTrigger,
	BtnLeftTrigger2,
	BtnRightTrigger,
	BtnRightTrigger2,
	BtnSelect,
	BtnStart,
	BtnMode,
	BtnLeftThumb,
	BtnRightThumb,
	BtnDPadUp,
	BtnDPadDown,
	BtnDPadLeft,
	BtnDPadRight,
}

var _BtnMap = map[uint8]string{
	'a': BtnSouth,
	'b': BtnEast,
	'c': BtnNorth,
	'd': BtnWest,
	'e': BtnC,
	'f': BtnZ,
	'g': BtnLeftTrigger,
	'h': BtnLeftTrigger2,
	'i': BtnRightTrigger,
	'j': BtnRightTrigger2,
	'k': BtnSelect,
	'l': BtnStart,
	'm': BtnMode,
	'n': BtnLeftThumb,
	'o': BtnRightThumb,
	'p': BtnDPadUp,
	'q': BtnDPadDown,
	'r': BtnDPadLeft,
	's': BtnDPadRight,
	't': BtnUnknown,
}

const (
	EvAxisChanged    string = "EvAxisChanged"
	EvButtonChanged         = "EvButtonChanged"
	EvButtonReleased        = "EvButtonReleased"
	EvButtonPressed         = "EvButtonPressed"
	EvButtonRepeated        = "EvButtonRepeated"
	EvConnected             = "EvConnected"
	EvDisconnected          = "EvDisconnected"
	EvDropped               = "EvDropped"
)

var ButtonEvents = []string{EvButtonChanged, EvButtonReleased, EvButtonPressed, EvButtonRepeated}

var EventTypeMap = map[uint8]string{
	'a': EvAxisChanged,
	'b': EvButtonChanged,
	'c': EvButtonReleased,
	'd': EvButtonPressed,
	'e': EvButtonRepeated,
	'f': EvConnected,
	'g': EvDisconnected,
	'h': EvDropped,
}

func genBtnAxisMap() map[uint8]string {
	mapping := map[uint8]string{}
	for k, v := range _AxisMap {
		AssignWithDuplicateCheck(mapping, k, v)
	}
	for k, v := range _BtnMap {
		AssignWithDuplicateCheck(mapping, k, v)
	}
	return mapping
}

var BtnAxisMap = genBtnAxisMap()
