package service

import "github.com/chrisvinsen/go-gin-currency/entity"

// Interface of Symbol Service
type SymbolService interface {
	Save([]entity.Symbol) []entity.Symbol
	FindAll() []entity.Symbol
}

// Struct of Symbol Service
type symbolService struct {
	symbols []entity.Symbol
}

// Initialize new symbol service
func NewSymbol() SymbolService {
	return &symbolService{
		symbols: []entity.Symbol{},
	}
}

// Method of Symbol Service to Save symbol data
func (service *symbolService) Save(symbols []entity.Symbol) []entity.Symbol {
	service.symbols = symbols
	return symbols
}

// Method of Symbol Service to Show all symbol data
func (service *symbolService) FindAll() []entity.Symbol {
	return service.symbols
}
