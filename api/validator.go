package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/tuda4/simple_bank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrencies(currency)
	}
	return false
}
