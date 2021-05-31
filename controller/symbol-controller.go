package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/chrisvinsen/go-gin-currency/entity"
	"github.com/chrisvinsen/go-gin-currency/service"
)

// Interface of Symbol Controller
type SymbolController interface {
	FindAll() []entity.Symbol
	Save(symbols []entity.Symbol) []entity.Symbol
	FetchExternalSymbolAPI() error
}

// Struct of Symbol Controller
type symbolController struct {
	service service.SymbolService
}

// Initialize new Symbol Controller
func NewSymbol(service service.SymbolService) SymbolController {
	return &symbolController{
		service: service,
	}
}

// Method of Symbol Controller to call FindAll services from Symbol Service in order to show all symbol data
func (c *symbolController) FindAll() []entity.Symbol {
	return c.service.FindAll()
}

// Method of Symbol Controller to call Save services from Symbol Service in order to save symbol data
func (c *symbolController) Save(symbols []entity.Symbol) []entity.Symbol {
	c.service.Save(symbols)
	return symbols
}

// Method of Symbol Controller to fetch symbol data from external API
func (c *symbolController) FetchExternalSymbolAPI() error {
	// Build The URL string
	API_KEY := "577392a42c95454cc4fef09b0e6040e9"
	URL := "http://api.exchangeratesapi.io/v1/symbols?access_key=" + API_KEY

	// Send HTTP Request GET
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response HTTP request
	var response entity.ResponseSymbolAPI
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	// Format response data to Symbol Entity Struct
	var symbols []entity.Symbol

	// Iterate struct Rates from response HTTP and append to Symbol Entity Struct
	v := reflect.ValueOf(response.Symbols)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		code := typeOfS.Field(i).Name
		name := v.Field(i).Interface().(string)
		symbols = append(symbols, entity.Symbol{
			Code: code,
			Name: name,
		})
	}

	fmt.Println("Successfully Fetch Symbols from External API")
	// Save formatted symbol data
	c.service.Save(symbols)
	return nil
}
