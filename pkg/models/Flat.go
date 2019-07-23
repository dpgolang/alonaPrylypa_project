package models

type Flat struct {
	Id      int     `json:"id"`
	Type    string  `json:"type"`
	Street  string  `json:"street"`
	Price   float64 `json:"price"`
	Square  float64 `json:"square"`
	Rooms   int     `json:"rooms"`
	Floor   int     `json:"floor"`
	Realtor string  `json:"realtor"`
}
type Customer struct {
	UserName string
	Email    string
	Password string
}
type User struct {
	Username      string
	Authenticated bool
}
type Realtor struct {
	Name  string
	Phone int
	Email string
}
type Privat struct{
	Date string `json:"date"`
	Bank string `json:"bank"`
	BaseCurrency string `json:"baseCurrency"`
	BaseCurrencyLit string `json:"baseCurrencyLit"`
	ExchangeRate []ExchangeRate `json:"exchangeRate"`
}
type ExchangeRate struct{
	BaseCurrency string `json:"baseCurrency"`
	Currency string `json:"currency"`
	SaleRateNB float64 `json:"saleRateNB"`
	PurchaseRateNB float64 `json:"purchaseRateNB"`
	SaleRate float64 `json:"saleRate"`
	PurchaseRate float64 `json:"purchaseRate"`
}
