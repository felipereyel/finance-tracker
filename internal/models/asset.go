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
	Id           string
	Name         string
	Type         string
	Wallet       string
	WalletName   string
	InitialPrice float32
	BuyDate      string
	LastPrice    float32
	LastDate     string
	SellDate     string
	Comment      string
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
	Id           string
	Created      string
	Updated      string
	Name         string  `form:"name"`
	Type         string  `form:"type"`
	Wallet       string  `form:"wallet"`
	Comment      string  `form:"comment"` // nullable
	InitialPrice float32 `form:"initial_price"`
	BuyDate      string  `form:"buy_date"`
	SellDate     string  `form:"sell_date"` // nullable
}

type AssetCreateDTO struct {
	Name         string  `form:"name"`
	Type         string  `form:"type"`
	Wallet       string  `form:"wallet"`
	Comment      string  `form:"comment"` // nullable
	InitialPrice float32 `form:"initial_price"`
	BuyDate      string  `form:"buy_date"`
	SellDate     string  `form:"sell_date"` // nullable
}

type AssetUpdateDTO struct {
	Comment  string `form:"comment"`   // nullable
	SellDate string `form:"sell_date"` // nullable
}

var AssetFields = []string{"id", "created", "updated", "name", "type", "wallet", "comment", "initial_price", "buy_date", "sell_date"}

var EmptyAsset = Asset{}

func CreateNewAsset(dto AssetCreateDTO) Asset {
	return Asset{
		Id:           GenerateId(),
		Created:      GenerateTimestamp(),
		Updated:      GenerateTimestamp(),
		Name:         dto.Name,
		Type:         dto.Type,
		Wallet:       dto.Wallet,
		Comment:      dto.Comment,
		InitialPrice: dto.InitialPrice,
		BuyDate:      dto.BuyDate,
		SellDate:     dto.SellDate,
	}
}
