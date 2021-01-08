package currency

type Currency interface {
	NewCurrency(name string) error
	Convert(convert CurrencyConvert) (float64, error)
	DeleteCurrency(currency Currency) error
	GetAllCurrencies() map[string]float64
}
