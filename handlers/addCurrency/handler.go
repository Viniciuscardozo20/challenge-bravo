package addCurrency

import (
	"challenge-bravo/api/controller"

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
	if request.Params["currency"] == "" {
		return httping.BadRequest(map[string]string{"currency": "the field currency is required"})
	}
	err := h.ctl.NewCurrency(request.Params["currency"])
	if err != nil {
		return httping.InternalServerError(map[string]string{"err": err.Error()})
	}
	return httping.OK(map[string]string{"success": "currency added"})
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
