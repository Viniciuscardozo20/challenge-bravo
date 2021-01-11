package convert

import (
	"challenge-bravo/api/controller"
	"strconv"

	httping "github.com/ednailson/httping-go"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	ctl controller.Controller
}

func NewHandler(ctl controller.Controller) *Handler {
	return &Handler{ctl: ctl}
}

func (h *Handler) Handle(request httping.HttpRequest) httping.IResponse {
	if request.Params["to"] == "" {
		return httping.BadRequest(map[string]string{"to": "the field to is required"})
	}
	if request.Params["from"] == "" {
		return httping.BadRequest(map[string]string{"to": "the field from is required"})
	}
	if request.Params["amount"] == "" {
		return httping.BadRequest(map[string]string{"to": "the field amount is required"})
	}
	amount, err := strconv.ParseFloat(request.Params["amount"], 64)
	if err != nil {
		return httping.BadRequest(map[string]string{"amount": "invalid type for amount"})
	}
	value, err := h.ctl.Convert(request.Params["to"], request.Params["from"], amount)
	if err != nil {
		return httping.InternalServerError(map[string]string{"err": err.Error()})
	}
	return httping.OK(value)
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
