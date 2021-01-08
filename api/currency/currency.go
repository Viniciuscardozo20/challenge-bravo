package currency

import (
	"challenge-bravo/api/currency/currency_updater"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const urlRates = `https://api.exchangeratesapi.io/latest?base=USD`

type CurrencyConvert struct {
	from   string
	to     string
	amount float64
}

func Init() error {
	err := currency_updater.CurrencyLive()
	if err != nil {
		return err
	}
	go func() {
		for {
			<-time.Tick(time.Minute * 2)
			currency_updater.CurrencyLive()
		}
	}()
	return nil
}

func NewCurrency(name string) error {
	name = strings.ToUpper(name)
	if _, ok := currency_updater.Currencies[name]; ok {
		return errors.New("This currency is already existis")
	}
	body, err := currency_updater.GetRates(urlRates)
	if err != nil {
		log.Fatal("Error to get currencies")
		return err
	}
	res := gjson.Get(string(body), fmt.Sprintf("rates.%s", name))
	if !res.Exists() {
		return errors.New("This currencys does not exists in our currency bases")
	}
	currency_updater.Currencies[name] = res.Float()
	return nil
}

func DeleteCurrency(name string) error {
	name = strings.ToUpper(name)
	if _, ok := currency_updater.Currencies[name]; ok {
		delete(currency_updater.Currencies, name)
		return nil
	}
	return errors.New("This currency does not exists")
}

func Convert(convert CurrencyConvert) (*float64, error) {
	convert.from = strings.ToUpper(convert.from)
	if _, ok := currency_updater.Currencies[convert.from]; ok {
		return nil, errors.New("This currency does not exists")
	}
	if _, ok := currency_updater.Currencies[convert.to]; ok {
		return nil, errors.New("This currency does not exists")
	}
	value := ((1 / currency_updater.Currencies[convert.from]) * currency_updater.Currencies[convert.to]) * convert.amount
	return &value, nil
}

func GetAllCurrencies() map[string]float64 {
	return currency_updater.Currencies
}
