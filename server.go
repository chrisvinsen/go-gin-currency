package main

import (
	"log"

	"github.com/chrisvinsen/go-gin-currency/controller"
	"github.com/chrisvinsen/go-gin-currency/service"
	"github.com/gin-gonic/gin"
)

// Initialize symbol (service and controller) and currency (service and controller)
var (
	symbolService    service.SymbolService       = service.NewSymbol()
	symbolController controller.SymbolController = controller.NewSymbol(symbolService)

	currencyService    service.CurrencyService       = service.NewCurrency()
	currencyController controller.CurrencyController = controller.NewCurrency(currencyService)
)

func main() {
	// Fetch Symbols and Currency Base Data from External API
	error_handler(symbolController.FetchExternalSymbolAPI())
	error_handler(currencyController.FetchExternalCurrencyAPI(symbolService))
	// Setting asynchronous scheduler every 1 minutes to update base data from External API
	go currencyController.SchedulerFetchExternalCurrencyAPI(symbolService)

	// Initialize Gin Gionic
	server := gin.Default()

	// Load static files (css and js) and HTML files
	server.Static("/css", "./templates/css")
	server.Static("/js", "./templates/js")
	server.LoadHTMLGlob("templates/*.html")

	// Route grouping for api
	apiRoutes := server.Group("/api")
	{
		// API to show all currency data
		apiRoutes.GET("/currency", func(ctx *gin.Context) {
			ctx.JSON(200, currencyController.FindAll())
		})

		// API to show details of currency data based on currency symbols
		apiRoutes.GET("/currency/:symbols", func(ctx *gin.Context) {
			ctx.JSON(200, currencyController.Find(ctx))
		})

		// API to show all symbols data
		apiRoutes.GET("/symbols", func(ctx *gin.Context) {
			ctx.JSON(200, symbolController.FindAll())
		})
	}

	// Route grouping for views
	viewRoutes := server.Group("/")
	{
		// View to show currency exchange views
		viewRoutes.GET("/", currencyController.ShowAll)
	}

	// Run server using port 8080
	server.Run(":8080")
}

func error_handler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
