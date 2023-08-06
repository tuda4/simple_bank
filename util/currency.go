package util

const (
	USD = "USD"
	UER = "UER"
	VND = "VND"
	GBP = "GBP"
)

func IsSupportedCurrencies(currency string) bool {
	switch currency {
	case USD, UER, VND, GBP:
		return true
	}
	return false
}
