package components

import (
	"fmt"
	"strconv"

	"github.com/NimbleMarkets/ntcharts/sparkline"
)

var (
	lengthOfGraph    int = 20
	temperatureQueue     = make([]float64, lengthOfGraph)
)

func StreamlineTemperatureChart(temperature string) string {
	coercion, err := strconv.Atoi(temperature)
	if err != nil {
		panic(err)
	}

	temperatureQueue = append(temperatureQueue, float64(coercion))

	temperatureQueue = temperatureQueue[1:]

	sparklineChart := sparkline.New(lengthOfGraph, 1)

	sparklineChart.PushAll(temperatureQueue)
	sparklineChart.Draw()

	return fmt.Sprintf("%v %s℃ ", sparklineChart.View(), temperature)
}
