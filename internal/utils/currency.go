package utils

import (
	"github.com/leekchan/accounting"
)

var ac = accounting.Accounting{Symbol: "R$", Precision: 2}

func FormatBRL(value float32) string {
	return ac.FormatMoney(value)
}
