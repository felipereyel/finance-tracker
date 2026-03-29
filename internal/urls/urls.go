package urls

import "fmt"

var Root = "/"

var DiscardURL = "/_discard"
var HealthzURL = "/_healthz"

func StaticsURL(path string) string {
	return fmt.Sprintf("/_statics/%s", path)
}

var AuthenticatedUrl = "/u"

// Query Parameters
var WalletQueryParam = "wallet"
var TagQueryParam = "tag"
var AggregationQueryParam = "aggregation"

func SummaryURLWithAggregation(aggregation string) string {
	return fmt.Sprintf("%s?%s=%s", SummaryURL, AggregationQueryParam, aggregation)
}

func SummaryChartURLWithAggregation(aggregation string) string {
	return fmt.Sprintf("%s?%s=%s", SummaryChartURL, AggregationQueryParam, aggregation)
}

var SummaryPath = "/summary"
var SummaryURL = AuthenticatedUrl + SummaryPath

var SummaryChartPath = "/summary-chart"
var SummaryChartURL = AuthenticatedUrl + SummaryChartPath

// History URLs
var HistoryPath = "/history"
var HistoryURL = AuthenticatedUrl + HistoryPath

var HistoryChartPath = "/history-chart"
var HistoryChartURL = AuthenticatedUrl + HistoryChartPath

func HistoryURLWithAggregation(aggregation string) string {
	return fmt.Sprintf("%s?%s=%s", HistoryURL, AggregationQueryParam, aggregation)
}

func HistoryChartURLWithAggregation(aggregation string) string {
	return fmt.Sprintf("%s?%s=%s", HistoryChartURL, AggregationQueryParam, aggregation)
}

var AssetsRedirectPath = "/assets-redirect"
var AssetsRedirectURL = AuthenticatedUrl + AssetsRedirectPath

var AssetsPopupPath = "/assets-popup"
var AssetsPopupURL = AuthenticatedUrl + AssetsPopupPath
