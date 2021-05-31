package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/chrisvinsen/go-gin-currency/entity"
	"github.com/chrisvinsen/go-gin-currency/service"
	"github.com/gin-gonic/gin"
)

// Interface of Currency Controller
type CurrencyController interface {
	FindAll() entity.Currency
	Find(ctx *gin.Context) gin.H
	Save(currency entity.Currency) entity.Currency
	ShowAll(ctx *gin.Context)
	FetchExternalCurrencyAPI(symbolService service.SymbolService) error
	SchedulerFetchExternalCurrencyAPI(symbolService service.SymbolService)
}

// Struct of Currency Controller
type currencyController struct {
	service service.CurrencyService
}

// Initialize new Currency Controller
func NewCurrency(service service.CurrencyService) CurrencyController {
	return &currencyController{
		service: service,
	}
}

// Method of Currency Controller to call FindAll services from Currency Service in order to show all currency data
func (c *currencyController) FindAll() entity.Currency {
	return c.service.FindAll()
}

// Method of Currency Controller to call Find services from Currency Service in order to show details of currency data based on symbols parameters
func (c *currencyController) Find(ctx *gin.Context) gin.H {
	requested_symbols := strings.ToUpper(ctx.Param("symbols"))

	return c.service.Find(requested_symbols)
}

// Method of Currency Controller to call Save services from Currency Service in order to save currency data
func (c *currencyController) Save(currency entity.Currency) entity.Currency {
	c.service.Save(currency)
	return currency
}

// Method of Currency Controller to Show All currency data in HTML format
func (c *currencyController) ShowAll(ctx *gin.Context) {
	currencies := c.service.FindAll()
	ctx.HTML(http.StatusOK, "index.html", currencies)
}

// Method of Currency Controller to fetch currency data from external API
func (c *currencyController) FetchExternalCurrencyAPI(symbolService service.SymbolService) error {
	// Build The URL string
	API_KEY := "577392a42c95454cc4fef09b0e6040e9"
	SUPPORTED_SYMBOLS := "USD,CAD,IDR,GBP,CHF,SGD,INR,MYR,JPY,KRW"
	URL := "http://api.exchangeratesapi.io/v1/latest?access_key=" + API_KEY + "&symbols=" + SUPPORTED_SYMBOLS

	// Send HTTP Request GET
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response HTTP request
	var response entity.ResponseCurrencyAPI
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	// Format response data to Currency Entity Struct
	currency := entity.Currency{
		Status: response.Success,
		Base:   response.Base,
		Rates:  []entity.Rate{},
	}

	// Get all symbols data
	symbolController := NewSymbol(symbolService)
	symbols := symbolController.FindAll()

	// Iterate struct Rates from response HTTP and append to Currency Entity Struct
	v := reflect.ValueOf(response.Rates)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		code := typeOfS.Field(i).Name
		rate := v.Field(i).Interface().(float64)
		name := GetSymbolName(symbols, code)

		if rate != 0 {
			currency.Rates = append(currency.Rates, entity.Rate{
				Code: code,
				Name: name,
				Rate: rate,
			})
		}
	}

	fmt.Println("Successfully Fetch Currency from External API")
	// Save formatted currency data
	c.service.Save(currency)
	return nil
}

// Method of Currency Controller to schedule fetching currency data from external API
func (c *currencyController) SchedulerFetchExternalCurrencyAPI(symbolService service.SymbolService) {
	// Schedule fetch external currency API every 60 seconds (1 minute)
	for range time.Tick(time.Second * 60) {
		err := c.FetchExternalCurrencyAPI(symbolService)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Function to get symbol name from his code
func GetSymbolName(symbols []entity.Symbol, code string) string {
	for i := range symbols {
		if symbols[i].Code == code {
			return symbols[i].Name
		}
	}
	return ""
}
