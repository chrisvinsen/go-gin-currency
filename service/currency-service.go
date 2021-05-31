package service

import (
	"fmt"

	"github.com/chrisvinsen/go-gin-currency/entity"
	"github.com/gin-gonic/gin"
)

// Interface of Currency Service
type CurrencyService interface {
	Save(entity.Currency) entity.Currency
	FindAll() entity.Currency
	Find(string) gin.H
}

// Struct of Currency Service
type currencyService struct {
	currency entity.Currency
}

// Initialize new currency service
func NewCurrency() CurrencyService {
	return &currencyService{
		currency: entity.Currency{},
	}
}

// Method of Currency Service to Save currency data
func (service *currencyService) Save(currency entity.Currency) entity.Currency {
	service.currency = currency
	return currency
}

// Method of Currency Service to Show all currency data
func (service *currencyService) FindAll() entity.Currency {
	return service.currency
}

// Method of Currency Service to Show details of currency data
func (service *currencyService) Find(symbols string) gin.H {
	for _, v := range service.currency.Rates {
		if v.Code == symbols {
			return gin.H{
				"status": true,
				"base":   service.currency.Base,
				"code":   v.Code,
				"name":   v.Name,
				"rate":   v.Rate,
			}
		}
	}

	return gin.H{
		"status":  false,
		"message": fmt.Sprintf("Currency with code %s not found.", symbols),
	}
}
