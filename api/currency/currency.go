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

type currency struct {
	updater *currency_updater.Updater
}

func Init() (Currency, error) {
	upd := currency_updater.NewUpdater()
	err := upd.CurrencyLive()
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			<-time.Tick(time.Minute * 2)
			upd.CurrencyLive()
		}
	}()
	return &currency{updater: upd}, nil
}

func (c *currency) NewCurrency(name string) error {
	name = strings.ToUpper(name)
	if _, ok := c.updater.Currencies[name]; ok {
		return errors.New("This currency is already existis")
	}
	body, err := c.updater.GetRates(urlRates)
	if err != nil {
		log.Fatal("Error to get currencies")
		return err
	}
	res := gjson.Get(string(body), fmt.Sprintf("rates.%s", name))
	if !res.Exists() {
		return errors.New("This currencys does not exists in our currency bases")
	}
	c.updater.Currencies[name] = res.Float()
	return nil
}

func (c *currency) DeleteCurrency(name string) error {
	name = strings.ToUpper(name)
	if _, ok := c.updater.Currencies[name]; ok {
		delete(c.updater.Currencies, name)
		return nil
	}
	return errors.New("This currency does not exists")
}

func (c *currency) Convert(convert CurrencyConvert) (*float64, error) {
	convert.from = strings.ToUpper(convert.from)
	if _, ok := c.updater.Currencies[convert.from]; ok {
		return nil, errors.New("This currency does not exists")
	}
	if _, ok := c.updater.Currencies[convert.to]; ok {
		return nil, errors.New("This currency does not exists")
	}
	value := ((1 / c.updater.Currencies[convert.from]) * c.updater.Currencies[convert.to]) * convert.amount
	return &value, nil
}

func (c *currency) GetAllCurrencies() map[string]float64 {
	return c.updater.Currencies
}
