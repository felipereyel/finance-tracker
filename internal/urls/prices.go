package urls

import "fmt"

// PRICE ID GROUP
var PricesPath = "/prices"

var PriceIdPathParam = "price_id"

func priceIdPath(priceId string) string {
	return fmt.Sprintf("%s/%s", PricesPath, priceId)
}

var PriceIdGroupPath = priceIdPath("{" + PriceIdPathParam + "}")

func PriceIdGroupURL(priceId string) string {
	return AuthenticatedUrl + priceIdPath(priceId) + Root
}
