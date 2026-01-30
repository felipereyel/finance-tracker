package models

type Price struct {
	Id       string
	Created  string
	Updated  string
	AssetId  string
	Value    float32
	LoggedAt string
	Gain     float32
	Comment  string // nullable
}

type PriceCreateDTO struct {
	Value    float32 `form:"value"`
	LoggedAt string  `form:"logged_at"`
	Comment  string  `form:"comment"` // nullable
}

type PriceUpdateDTO struct {
	Comment string `form:"comment"` // nullable
}

var EmptyPrice = Price{}

func CreateNewPrice(assetId string, dto PriceCreateDTO) Price {
	return Price{
		Id:       GenerateId(),
		Created:  GenerateTimestamp(),
		Updated:  GenerateTimestamp(),
		AssetId:  assetId,
		Value:    dto.Value,
		LoggedAt: dto.LoggedAt,
		Comment:  dto.Comment,
	}
}
