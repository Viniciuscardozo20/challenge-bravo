package app

import (
	"challenge-bravo/api/controller"
	"challenge-bravo/api/currency"
	"challenge-bravo/handlers/addCurrency"
	"challenge-bravo/handlers/convert"
	"challenge-bravo/handlers/deleteCurrency"
	"challenge-bravo/handlers/getCurrencies"

	httping "github.com/ednailson/httping-go"
	log "github.com/sirupsen/logrus"
)

type App struct {
	server    httping.IServer
	closeFunc httping.ServerCloseFunc
}

func LoadApp() (*App, error) {
	var app App
	curr, err := currency.Init()
	if err != nil {
		return nil, err
	}
	controller := controller.NewController(curr)
	app.server = loadServer(controller)
	return &app, nil
}

func (a *App) Run() <-chan error {
	closeFunc, chErr := a.server.RunServer()
	a.closeFunc = closeFunc
	return chErr
}

func (a *App) Close() {
	err := a.closeFunc()
	if err != nil {
		log.WithField("error", err.Error()).Errorf("failed to close func")
	}
}

func loadServer(ctrl *controller.Controller) httping.IServer {
	server := httping.NewHttpServer("", 8082)
	addCurrencyHandler := addCurrency.NewHandler(*ctrl)
	server.NewRoute(nil, "/v1/addCurrency/:currency").POST(addCurrencyHandler.Handle)
	convertHandler := convert.NewHandler(*ctrl)
	server.NewRoute(nil, "/v1/currency").POST(convertHandler.Handle)
	deleteCurrencyHandler := deleteCurrency.NewHandler(*ctrl)
	server.NewRoute(nil, "/v1/deleteCurrency/:currency").POST(deleteCurrencyHandler.Handle)
	getCurrenciesHandler := getCurrencies.NewHandler(*ctrl)
	server.NewRoute(nil, "/v1/getCurrencies").POST(getCurrenciesHandler.Handle)
	return server
}
