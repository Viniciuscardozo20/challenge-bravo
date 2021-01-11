package getCurrencies

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
	currencies := h.ctl.GetAllCurrencies()
	return httping.OK(currencies)
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}
