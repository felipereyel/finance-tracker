package models

// Type:  "fii" | "federal_bond" | "cdb" | "hedge_fund" | "stock" | "other"

type Asset struct {
	Id           string  `json:"id"`
	Created      string  `json:"created"`
	Updated      string  `json:"updated"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Wallet       string  `json:"wallet"`
	Comment      string  `json:"comment"` // nullable
	InitialPrice float32 `json:"initial_price"`
	Bought       string  `json:"buy_date"`
	Sold         string  `json:"sell_date"` // nullable
}

var AssetFields = []string{"id", "created", "updated", "name", "type", "wallet", "comment", "initial_price", "buy_date", "sell_date"}

var EmptyAsset = Asset{}
