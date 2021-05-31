package entity

type Currency struct {
	Status bool   `json:"status"`
	Base   string `json:"base"`
	Rates  []Rate `json:"rates"`
}

type Rate struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}
