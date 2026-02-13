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
	page := components.NewPage().SetAssetsHost(urls.StaticsURL("assets/"))

	switch summary.Aggregation {
	case "wallet":
		walletPie := buildWalletPieChart(summary)
		page.AddCharts(walletPie)
	case "type":
		typePie := buildTypePieChart(summary)
		page.AddCharts(typePie)
	case "total":
		totalChart := buildTotalChart(summary)
		page.AddCharts(totalChart)
	default:
		// Default to showing both charts
		walletPie := buildWalletPieChart(summary)
		typePie := buildTypePieChart(summary)
		page.AddCharts(walletPie, typePie)
	}

	return page.Render(w)
}

func formatCurrency(value float32) string {
	return fmt.Sprintf("R$%.2f", value)
}

func buildWalletPieChart(summary models.Summary) *charts.Pie {
	walletMap := make(map[string]float32)
	for _, wallet := range summary.Wallets {
		walletMap[wallet[1]] = 0.0
	}

	for _, asset := range summary.Aggregates {
		if _, ok := walletMap[asset.WalletName]; ok {
			walletMap[asset.WalletName] += asset.LastPrice
		}
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
			Height:          "400px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Portifolio By Wallet",
			Subtitle: "Total: " + formatCurrency(summary.Total),
			TitleStyle: &opts.TextStyle{
				Color: "#E2E8F0",
			},
			SubtitleStyle: &opts.TextStyle{
				Color: "#94A3B8",
			},
		}),
	)

	walletPie.AddSeries("area", walletItems,
		charts.WithLabelOpts(opts.Label{
			Show:      opts.Bool(true),
			Formatter: "{b}: R${c}",
		}),
	)

	return walletPie
}

func buildTypePieChart(summary models.Summary) *charts.Pie {
	typeMap := make(map[string]float32)
	for _, assetType := range summary.AssetTypes {
		typeMap[assetType[0]] = 0.0
	}

	for _, asset := range summary.Aggregates {
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

	typePie := charts.NewPie()
	typePie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(false),
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme:           types.ThemeWonderland,
			BackgroundColor: "#0F172A",
			Height:          "400px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Portifolio By Asset Type",
			Subtitle: "Total: " + formatCurrency(summary.Total),
			TitleStyle: &opts.TextStyle{
				Color: "#E2E8F0",
			},
			SubtitleStyle: &opts.TextStyle{
				Color: "#94A3B8",
			},
		}),
	)

	typePie.AddSeries("radius", typeItems,
		charts.WithLabelOpts(opts.Label{
			Show:      opts.Bool(true),
			Formatter: "{b}: R${c}",
		}),
	)

	return typePie
}

func buildTotalChart(summary models.Summary) *charts.Pie {
	totalPie := charts.NewPie()
	totalPie.SetGlobalOptions(
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(false),
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme:           types.ThemeWonderland,
			BackgroundColor: "#0F172A",
			Height:          "400px",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Total Portfolio Value",
			Subtitle: "Total: " + formatCurrency(summary.Total),
			TitleStyle: &opts.TextStyle{
				Color: "#E2E8F0",
			},
			SubtitleStyle: &opts.TextStyle{
				Color: "#94A3B8",
			},
		}),
	)

	totalPie.AddSeries("total", []opts.PieData{{Name: "Total", Value: summary.Total}},
		charts.WithLabelOpts(opts.Label{
			Show:      opts.Bool(true),
			Formatter: "{b}: R${c}",
		}),
	)

	return totalPie
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

func HistoryChart(history models.History, w io.Writer) error {
	page := components.NewPage().SetAssetsHost(urls.StaticsURL("assets/"))
	line := charts.NewLine()

	// Calculate min/max for axis
	var maxValue float32 = 0.0
	var minValue float32 = 0.0
	minDate := time.Date(2100, time.January, 1, 0, 0, 0, 0, time.Local)
	maxDate := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)

	// If no series or empty, use dummy data
	if len(history.Series) == 0 {
		history.Series = generateDummySeries()
	}

	for _, series := range history.Series {
		items := make([]opts.LineData, 0)
		for _, point := range series.Points {
			t, err := time.Parse(DATE_LAYOUT, point.Date[:DATE_LAYOUT_LEN])
			if err != nil {
				fmt.Println("Error parsing date:", err)
				continue
			}

			if point.Value > maxValue || maxValue == 0.0 {
				maxValue = point.Value
			}
			if point.Value < minValue || minValue == 0.0 {
				minValue = point.Value
			}
			if t.Before(minDate) {
				minDate = t
			}
			if t.After(maxDate) {
				maxDate = t
			}

			items = append(items, opts.LineData{
				Value: []interface{}{t, point.Value},
			})
		}
		line.AddSeries(series.Name, items,
			charts.WithLineChartOpts(opts.LineChart{
				Symbol: "none",
			}),
		)
	}

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
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    opts.Bool(true),
			Trigger: "axis",
			Formatter: opts.FuncOpts(`
				function (info) {
					var result = '<b>' + info[0].axisValueLabel + '</b><br/>';
					for (var i = 0; i < info.length; i++) {
						var value = info[i].value[1].toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' });
						result += info[i].marker + ' ' + info[i].seriesName + ': ' + value + '<br/>';
					}
					return result;
				}
			`),
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: opts.Bool(true),
			Top:  "5%",
			TextStyle: &opts.TextStyle{
				Color: "#E2E8F0",
			},
		}),
	)

	page.AddCharts(line)
	return page.Render(w)
}

func generateDummySeries() []models.HistorySeries {
	series := []models.HistorySeries{
		{
			Name:   "Wallet 1",
			Points: generateDummyPoints(30, 1000, 5000),
		},
		{
			Name:   "Wallet 2",
			Points: generateDummyPoints(30, 2000, 8000),
		},
		{
			Name:   "Wallet 3",
			Points: generateDummyPoints(30, 1500, 6000),
		},
	}
	return series
}

func generateDummyPoints(count int, minVal, maxVal float32) []models.HistoryDataPoint {
	points := make([]models.HistoryDataPoint, count)
	baseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)

	for i := 0; i < count; i++ {
		date := baseDate.AddDate(0, 0, i)
		value := minVal + float32(i)*((maxVal-minVal)/float32(count))
		points[i] = models.HistoryDataPoint{
			Date:  date.Format(DATE_LAYOUT),
			Value: value,
		}
	}
	return points
}
