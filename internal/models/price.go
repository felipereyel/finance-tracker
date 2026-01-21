package models

type Price struct {
	Id      string
	Created string
	Updated string
	AssetId string
	Value   float32
	Logged  string
	Gain    float32
	Comment string // nullable
}

type PriceCreateDTO struct {
	AssetId string  `form:"asset_id"`
	Value   float32 `form:"value"`
	Logged  string  `form:"logged"`
	Comment string  `form:"comment"` // nullable
}

var EmptyPrice = Price{}

func CreateNewPrice(dto PriceCreateDTO) Price {
	return Price{
		Id:      GenerateId(),
		Created: GenerateTimestamp(),
		Updated: GenerateTimestamp(),
		AssetId: dto.AssetId,
		Value:   dto.Value,
		Logged:  dto.Logged,
		Comment: dto.Comment,
	}
}
