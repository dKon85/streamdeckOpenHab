package openhab

import (
	"strconv"
	"strings"
)

const (
	TempWarmLimit = 21.00
	TempColdLimit = 18.00
	TempOffLimit = 4.3
)

func ConvertTemperatureToFloat(tempState string) float32 {
	split := strings.Split(tempState, " ")
	if s, err := strconv.ParseFloat(split[0], 32); err == nil {
		return float32(s)
	}
	return 0.0
}

func TempToName(temp float32) string{
	if temp >= TempWarmLimit {
		return "warm"
	} else if temp >= TempColdLimit {
		return "cold"
	} else {
		return "off"
	}
}

func DetermineNextTempState( curState string) string{
	switch curState {
	case "warm":
		return "cold"
	case "cold":
		return "warm"
	case "off":
		return "cold"
	}

	return "cold"
}
