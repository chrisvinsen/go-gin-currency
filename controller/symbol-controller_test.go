package controller

import (
	"testing"

	"github.com/chrisvinsen/go-gin-currency/entity"
	"github.com/chrisvinsen/go-gin-currency/service"

	"github.com/stretchr/testify/assert"
)

// Initialize Symbol Controller and Symbol Entities
var symbolControllerTest SymbolController = NewSymbol(service.NewSymbol())
var symbolTest []entity.Symbol = []entity.Symbol{
	{
		Code: "IDR",
		Name: "Indonesian Rupiah",
	},
	{
		Code: "SGD",
		Name: "Singapore Dollar",
	},
}

// Unit test of methods Save Symbol in Symbol Controller
func TestSaveSymbol(t *testing.T) {
	symbolSaved := symbolControllerTest.Save(symbolTest)
	// Compare the result with the input
	assert.Equal(t, symbolTest, symbolSaved)
}

// Unit test of methods Find All Symbol in Symbol Controller
func TestFindAllSymbol(t *testing.T) {
	symbolFound := symbolControllerTest.FindAll()
	// Compare the result with the input
	assert.Equal(t, symbolTest, symbolFound)
}

// Unit test of methods Fetch External Symbol API in Symbol Controller
func TestFetchExternalSymbolAPI(t *testing.T) {
	symbolController := NewSymbol(service.NewSymbol())
	// Check whether the result is an error or not
	assert.Nil(t, symbolController.FetchExternalSymbolAPI())
}
