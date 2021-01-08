package controller

import "challenge-bravo/api/currency"

type Controller struct {
	curr currency.Currency
}

func NewController(curr currency.Currency) *Controller {
	return &Controller{
		curr: curr,
	}
}
