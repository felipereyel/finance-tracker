package urls

import "fmt"

var AssetsPath = "/assets"
var AssetsURL = AuthenticatedUrl + AssetsPath

func AssetsURLWithQuey(walletFilter string, typefilter string) string {
	return fmt.Sprintf("%s?wallet=%s&type=%s", AssetsURL, walletFilter, typefilter)
}

// ASSET ID GROUP
var AssetIdPathParam = "asset_id"

func AssetIdPath(assetId string) string {
	return fmt.Sprintf("%s/%s", AssetsPath, assetId)
}

var AssetIdGroupPath = AssetIdPath("{" + AssetIdPathParam + "}")

func assetIdURL(assetId string, path string) string {
	return AuthenticatedUrl + AssetIdPath(assetId) + path
}

func AssetIdGroupURL(assetId string) string {
	return assetIdURL(assetId, Root)
}

var AssetIdPricesPath = "/prices"

func AssetIdPricesURL(assetId string) string {
	return assetIdURL(assetId, AssetIdPricesPath)
}

var AssetIdPricesChartPath = "/prices-chart"

func AssetIdPricesChartURL(assetId string) string {
	return assetIdURL(assetId, AssetIdPricesChartPath)
}

var AssetIdPricesPopupPath = "/prices-popup"

func AssetIdPricesPopupURL(assetId string) string {
	return assetIdURL(assetId, AssetIdPricesPopupPath)
}
