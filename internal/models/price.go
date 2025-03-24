package models

type Price struct {
	Id      string  `json:"id"`
	Created string  `json:"created"`
	Updated string  `json:"updated"`
	AssetId string  `json:"asset_id"`
	Value   float32 `json:"value"`
	Logged  string  `json:"logged_at"`
	Gain    float32 `json:"gain"`
	Comment string  `json:"comment"` // nullable
}

var EmptyPrice = Price{}
