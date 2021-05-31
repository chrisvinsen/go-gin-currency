package entity

type ResponseCurrencyAPI struct {
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     struct {
		USD float64 `json:"USD"`
		CAD float64 `json:"CAD"`
		IDR float64 `json:"IDR"`
		GBP float64 `json:"GBP"`
		CHF float64 `json:"CHF"`
		SGD float64 `json:"SGD"`
		INR float64 `json:"INR"`
		MYR float64 `json:"MYR"`
		JPY float64 `json:"JPY"`
		KRW float64 `json:"KRW"`
	} `json:"rates"`
}

type ResponseSymbolAPI struct {
	Symbols struct {
		USD string `json:"USD"`
		CAD string `json:"CAD"`
		IDR string `json:"IDR"`
		GBP string `json:"GBP"`
		CHF string `json:"CHF"`
		SGD string `json:"SGD"`
		INR string `json:"INR"`
		MYR string `json:"MYR"`
		JPY string `json:"JPY"`
		KRW string `json:"KRW"`
	} `json:"symbols"`
}
