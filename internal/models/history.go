package models

type History struct {
	Aggregation string
	Series      []HistorySeries
	Wallets     [][]string
	AssetTypes  [][]string
}

type HistorySeries struct {
	Name   string
	Points []HistoryDataPoint
}

type HistoryDataPoint struct {
	Date  string
	Value float32
}

type AssetPriceHistory struct {
	AssetId    string
	AssetName  string
	AssetType  string
	WalletId   string
	WalletName string
	SellDate   string
	Prices     []Price
}

var EmptyHistory = History{}
