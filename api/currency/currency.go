package currency

type currency struct {
	name  string
	value float32
}

type currencyConvert struct {
	from   string
	to     string
	amount float32
}

type currencies struct {
	currencies []currency
}

func NewCurrency(name string) (*currency, error) {
	// var currency = currency{
	// 	name:  strings.ToUpper(name),
	// 	value: 0,
	// }
	return nil, nil
}
