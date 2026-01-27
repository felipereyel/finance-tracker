package urls

import "fmt"

var Root = "/"

var DiscardURL = "/_discard"
var HealthzURL = "/_healthz"

func StaticsURL(path string) string {
	return fmt.Sprintf("/_statics/%s", path)
}

var AuthenticatedUrl = "/u"

var SummaryPath = "/summary"
var SummaryURL = AuthenticatedUrl + SummaryPath

var SummaryChartPath = "/summary-chart"
var SummaryChartURL = AuthenticatedUrl + SummaryChartPath

var AssetsRedirectPath = "/assets-redirect"
var AssetsRedirectURL = AuthenticatedUrl + AssetsRedirectPath

var AssetsPopupPath = "/assets-popup"
var AssetsPopupURL = AuthenticatedUrl + AssetsPopupPath
