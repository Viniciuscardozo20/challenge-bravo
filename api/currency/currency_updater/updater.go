package currency_updater

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

const urlRates = `https://api.exchangeratesapi.io/latest?base=USD`
const btcEthUrl = `https://min-api.cryptocompare.com/data/price?fsym=USD&tsyms=ETH,BTC`

var Currencies = map[string]float64{
	"USD": 0,
	"BRL": 0,
	"EUR": 0,
	"BTC": 0,
	"ETH": 0,
}

func CurrencyLive() error {
	body, err := GetRates(urlRates)
	if err != nil {
		log.Fatal("Error to get currencies")
		return err
	}
	btcEth, err := GetRates(btcEthUrl)
	if err != nil {
		log.Fatal("Error to get BTC and ETH currencies")
		return err
	}
	for currency, _ := range Currencies {
		if currency == "BTC" || currency == "ETH" {
			Currencies[currency] = gjson.Get(string(btcEth), fmt.Sprintf("%s", currency)).Float()
			continue
		}
		Currencies[currency] = gjson.Get(string(body), fmt.Sprintf("rates.%s", currency)).Float()
	}
	return nil
}

func GetRates(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, http.ErrServerClosed
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, http.ErrServerClosed
	}
	return body, nil
}
