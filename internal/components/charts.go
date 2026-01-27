package components

import (
	"encoding/json"
	"fintracker/internal/models"
	"fintracker/internal/urls"
	"fmt"
	"io"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func floorByMagnitude(value float32, magnitude float32) float32 {
	return float32(int(value/magnitude)) * magnitude
}

func ceilByMagnitude(value float32, magnitude float32) float32 {
	return float32(int(value/magnitude)+1) * magnitude
}

func SummaryChart(summary models.Summary, w io.Writer) error {
	walletMap := make(map[string]float32)
	for _, wallet := range summary.Wallets {
		walletMap[wallet[1]] = 0.0
	}

	typeMap := make(map[string]float32)
	for _, assetType := range summary.AssetTypes {
		typeMap[assetType[0]] = 0.0
	}

	for _, asset := range summary.Aggregates {
		if _, ok := walletMap[asset.WalletName]; ok {
			walletMap[asset.WalletName] += asset.LastPrice
		}

		if _, ok := typeMap[asset.Type]; ok {
			typeMap[asset.Type] += asset.LastPrice
		}
	}

	typeItems := make([]opts.PieData, 0)
	for assetType, value := range typeMap {
		if value == 0.0 {
			continue
		}

		typeItems = append(typeItems, opts.PieData{Name: models.GetLabelForType(assetType), Value: value})
	}

	walletItems := make([]opts.PieData, 0)
	for wallet, value := range walletMap {
		if value == 0.0 {
			continue
		}

		walletItems = append(walletItems, opts.PieData{Name: wallet, Value: value})
	}

	walletPie := charts.NewPie()
	walletPie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(false),
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme:           types.ThemeWonderland,
			BackgroundColor: "#0F172A",
			Height:          "200px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Wallet Summary",
		}),
	)

	walletPie.AddSeries("area", walletItems,
		charts.WithLabelOpts(opts.Label{
			Show:      opts.Bool(true),
			Formatter: "{b}: R${c}",
		}),
	)

	typePie := charts.NewPie()
	typePie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(false),
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme:           types.ThemeWonderland,
			BackgroundColor: "#0F172A",
			Height:          "200px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title: "Asset Type Summary",
		}),
	)

	typePie.AddSeries("radius", typeItems,
		charts.WithLabelOpts(opts.Label{
			Show:      opts.Bool(true),
			Formatter: "{b}: R${c}",
		}),
	)

	return components.NewPage().AddCharts(walletPie, typePie).SetAssetsHost(urls.StaticsURL("assets/")).Render(w)
}

var DATE_LAYOUT = "2006-01-02"
var DATE_LAYOUT_LEN = len(DATE_LAYOUT)

func PriceChart(prices []models.Price, w io.Writer) error {
	items := make([]opts.LineData, 0)

	minDate := time.Date(2100, time.January, 1, 0, 0, 0, 0, time.Local)
	maxDate := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)

	var maxValue float32 = 0.0
	var minValue float32 = 0.0

	for _, price := range prices {
		t, err := time.Parse(DATE_LAYOUT, price.LoggedAt[:DATE_LAYOUT_LEN])
		if err != nil {
			fmt.Println("Error parsing date:", err)
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

		data := struct {
			Id      string `json:"name"`
			Comment string `json:"comment"`
		}{
			Id:      price.Id,
			Comment: price.Comment,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshaling json:", err)
			continue
		}

		items = append(items, opts.LineData{Value: []interface{}{t, price.Value, string(jsonData)}})
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
			Show:      opts.Bool(true),
			Trigger:   "axis",
			Formatter: opts.FuncOpts(ToolTipFormatter),
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(false),
		}),
	)

	line.AddSeries("Price", items)
	return components.NewPage().AddCharts(line).SetAssetsHost(urls.StaticsURL("assets/")).Render(w)
}

var ToolTipFormatter = `
function (info) {
	var jsonData = info[0].value[2];
	var data = JSON.parse(jsonData);
	var id = data.id;
	var comment = data.comment;
	var value = info[0].value[1].toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });

	var tooltip = '<b>' + value + '</b>';
	tooltip += comment ? ('<br/>' + comment) : '';

	return tooltip;
}
`
