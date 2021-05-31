package controller

import (
	"testing"

	"github.com/chrisvinsen/go-gin-currency/entity"
	"github.com/chrisvinsen/go-gin-currency/service"

	"github.com/stretchr/testify/assert"
)

// Initialize Currency Controller and Currency Entities
var currencyControllerTest CurrencyController = NewCurrency(service.NewCurrency())
var currencyTest entity.Currency = entity.Currency{
	Status: true,
	Base:   "USD",
	Rates: []entity.Rate{
		{
			Code: "IDR",
			Name: "Indonesian Rupiah",
			Rate: 14000,
		},
		{
			Code: "SGD",
			Name: "Singapore Dollar",
			Rate: 5000,
		},
	},
}

// Unit test of methods Save Currency in Currency Controller
func TestSaveCurrency(t *testing.T) {
	currencySaved := currencyControllerTest.Save(currencyTest)
	// Compare the result with the input
	assert.Equal(t, currencyTest, currencySaved)
}

// Unit test of methods Find All Currency in Currency Controller
func TestFindAllCurrency(t *testing.T) {
	currencyFound := currencyControllerTest.FindAll()
	// Compare the result with the input
	assert.Equal(t, currencyTest, currencyFound)
}

// Unit test of methods Fetch External Currency API in Currency Controller
func TestFetchExternalCurrencyAPI(t *testing.T) {
	currencyController := NewCurrency(service.NewCurrency())
	// Check whether the result is an error or not
	assert.Nil(t, currencyController.FetchExternalCurrencyAPI(service.NewSymbol()))
}

// Unit test of functions in Currency Controller
func TestGetSymbolName(t *testing.T) {
	// Initialize list of symbols
	symbols := []entity.Symbol{
		{
			Code: "USD",
			Name: "United States Dollar",
		},
		{
			Code: "IDR",
			Name: "Indonesian Rupiah",
		},
		{
			Code: "SGD",
			Name: "Singapore Dollar",
		},
	}

	// Check the symbols name according to the code
	assert.Equal(t, "Indonesian Rupiah", GetSymbolName(symbols, "IDR"))
}
