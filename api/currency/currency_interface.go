package currency

type Currency interface {
	NewCurrency(name string) (*currency, error)
	CurrencyConvert(convert currencyConvert) (float32, error)
	DeleteCurrency(currency Currency) error
	GetCurrency(name string) (*currency, error)
}
