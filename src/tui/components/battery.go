package components

import (
	"fmt"

	"strconv"
)

func BatteryInformationComponent(capacity string, chargingStatus string) string {
	var (
		indicator = ""
	)

	coercion, err := strconv.Atoi(capacity)
	if err != nil {
		panic(err)
	}

	if coercion <= 10 {
		indicator = "󰂃 "
	} else if coercion <= 20 {
		indicator = "󰁻 "
	} else if coercion <= 40 {
		indicator = "󰁽 "
	} else if coercion <= 60 {
		indicator = "󰁿 "
	} else if coercion <= 80 {
		indicator = "󰂁 "
	} else if coercion <= 100 {
		indicator = "󰁹 "
	}

	switch chargingStatus {
	case "Charging":
		indicator += "󱐋 "
	case "Not charging":
		indicator += " "
	}

	return fmt.Sprintf("%d%% %s", coercion, indicator)
}
