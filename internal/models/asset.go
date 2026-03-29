package models

import (
	"encoding/json"
	"strings"
)

// ParseTags parses a PocketBase JSON-encoded tag string (e.g. `["A","B"]`) into a []string.
// Returns an empty slice for empty or invalid values.
func ParseTags(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" || raw == "[]" || raw == "null" {
		return []string{}
	}
	var tags []string
	if err := json.Unmarshal([]byte(raw), &tags); err != nil {
		return []string{}
	}
	return tags
}

// EncodeTags encodes a []string into the JSON format PocketBase expects.
func EncodeTags(tags []string) string {
	if len(tags) == 0 {
		return "[]"
	}
	b, err := json.Marshal(tags)
	if err != nil {
		return "[]"
	}
	return string(b)
}

// TagContains reports whether the raw JSON tag string contains the given tag.
func TagContains(raw string, tag string) bool {
	for _, t := range ParseTags(raw) {
		if t == tag {
			return true
		}
	}
	return false
}

// JoinTags returns a comma-separated string of tags, or a fallback if empty.
func JoinTags(raw string, empty string) string {
	tags := ParseTags(raw)
	if len(tags) == 0 {
		return empty
	}
	return strings.Join(tags, ", ")
}

var AssetTags = []string{
	"Renda Fixa",
	"Renda Variavel",
	"Fundos",
	"Acoes",
	"FII",
	"Tesouro",
	"CDB",
	"Reserva Emergencia",
	"Previdencia",
}

type AssetAggregate struct {
	Id           string
	Name         string
	Tag          string
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

	Wallets        [][]string
	SelectedWallet string

	AssetTags   []string
	SelectedTag string

	Aggregation string
}

type NewAssetOptions struct {
	AssetTags []string
	Wallets   [][]string
}

type Asset struct {
	Id           string
	Created      string
	Updated      string
	Name         string
	Tag          string
	Wallet       string
	Comment      string // nullable
	InitialPrice float32
	BuyDate      string
	SellDate     string // nullable
}

type AssetCreateDTO struct {
	Name         string   `form:"name"`
	Tags         []string `form:"tags"` // multi-select; encoded to JSON before storing
	Wallet       string   `form:"wallet"`
	Comment      string   `form:"comment"` // nullable
	InitialPrice float32  `form:"initial_price"`
	BuyDate      string   `form:"buy_date"`
	SellDate     string   `form:"sell_date"` // nullable
}

type AssetUpdateDTO struct {
	Comment  string `form:"comment"`   // nullable
	SellDate string `form:"sell_date"` // nullable
}

var AssetFields = []string{"id", "created", "updated", "name", "wallet", "comment", "initial_price", "buy_date", "sell_date"}

var EmptyAsset = Asset{}

var EmptySummary = Summary{}

var EmptyNewAssetOptions = NewAssetOptions{}

var EmptyAssetAggregate = AssetAggregate{}

func CreateNewAsset(dto AssetCreateDTO) Asset {
	return Asset{
		Id:           GenerateId(),
		Created:      GenerateTimestamp(),
		Updated:      GenerateTimestamp(),
		Name:         dto.Name,
		Tag:          EncodeTags(dto.Tags),
		Wallet:       dto.Wallet,
		Comment:      dto.Comment,
		InitialPrice: dto.InitialPrice,
		BuyDate:      dto.BuyDate,
		SellDate:     dto.SellDate,
	}
}
