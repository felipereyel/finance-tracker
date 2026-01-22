package components

import (
	"fintracker/internal/models"
	"io"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func floorByMagnitude(value float32, magnitude float32) float32 {
	return float32(int(value/magnitude)) * magnitude
}

func ceilByMagnitude(value float32, magnitude float32) float32 {
	return float32(int(value/magnitude)+1) * magnitude
}

func PriceChart(prices []models.Price, w io.Writer) error {
	items := make([]opts.LineData, 0)

	minDate := time.Date(2100, time.January, 1, 0, 0, 0, 0, time.Local)
	maxDate := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)

	var maxValue float32 = 0.0
	var minValue float32 = 0.0

	for _, price := range prices {
		t, err := time.Parse("2006-01-02 00:00:00.000Z", price.LoggedAt)
		if err != nil {
			continue
		}

		if price.Value > maxValue || maxValue == 0.0 {
			maxValue = price.Value
		}

		if price.Value < minValue || minValue == 0.0 {
			minValue = price.Value
		}

		if t.Before(minDate) {
			minDate = t
		}

		if t.After(maxDate) {
			maxDate = t
		}

		items = append(items, opts.LineData{Value: []interface{}{t, price.Value, price.Id}})
	}

	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme:           types.ThemeWonderland,
			BackgroundColor: "#0F172A",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Min: floorByMagnitude(minValue, 100),
			Max: ceilByMagnitude(maxValue, 100),
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Type: "time",
			Min:  minDate,
		}),
		charts.WithTooltipOpts(opts.Tooltip{ // Potential to string format tooltip here
			Show:      true,
			Trigger:   "axis",
			Formatter: opts.FuncOpts(ToolTipFormatter),
		}),
	)

	line.AddSeries("Price", items)
	return line.Render(w)
}

var ToolTipFormatter = `
function (info) {
	var id = info[0].value[2];
	var value = info[0].value[1].toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

	return '<a href="/prices/' + id + '">' + value + '</a>';
}
`
