package models

// Type:  "fii" | "federal_bond" | "cdb" | "hedge_fund" | "stock" | "other"
var AssetTypes = [][]string{
	{"fii", "FII"},
	{"federal_bond", "Federal Bond"},
	{"cdb", "CDB"},
	{"hedge_fund", "Hedge Fund"},
	{"stock", "Stock"},
	{"other", "Other"},
}

func GetLabelForType(assetType string) string {
	for _, at := range AssetTypes {
		if at[0] == assetType {
			return at[1]
		}
	}

	return "Unknown"
}

type AssetAggregate struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Wallet       string  `json:"wallet"`
	WalletName   string  `json:"wallet_name"`
	InitialPrice float32 `json:"initial_price"`
	BuyDate      string  `json:"buy_date"`
	LastPrice    float32 `json:"last_price"`
	LastDate     string  `json:"last_date"`
	SellDate     string  `json:"sell_date"`
	Comment      string  `json:"comment"`
}

type Summary struct {
	Total      float32
	Aggregates []AssetAggregate

	AssetTypes   [][]string
	SelectedType string

	Wallets        [][]string
	SelectedWallet string
}

type NewAssetSummary struct {
	AssetTypes [][]string
	Wallets    [][]string
}

type Asset struct {
	Id           string  `json:"id"`
	Created      string  `json:"created"`
	Updated      string  `json:"updated"`
	Name         string  `json:"name" form:"name"`
	Type         string  `json:"type" form:"type"`
	Wallet       string  `json:"wallet" form:"wallet"`
	Comment      string  `json:"comment" form:"comment"` // nullable
	InitialPrice float32 `json:"initial_price" form:"initial_price"`
	BuyDate      string  `json:"buy_date" form:"buy_date"`
	SellDate     string  `json:"sell_date"` // nullable
}

var AssetFields = []string{"id", "created", "updated", "name", "type", "wallet", "comment", "initial_price", "buy_date", "sell_date"}

var EmptyAsset = Asset{}
