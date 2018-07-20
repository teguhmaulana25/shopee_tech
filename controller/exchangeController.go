package controller

import (
	"net/http"

	"github.com/teguhmaulana25/shopee_tech/model"
)

var exchange = model.NewExchangeModel()

type ExchangeController struct{}

func NewExchangeController() ExchangeController {
	return ExchangeController{}
}

func (ExchangeController) Index(w http.ResponseWriter, r *http.Request) {
	response := res{
		Code:    200,
		Message: "Success",
		Data:    "",
	}

	renderJSON(w, response, http.StatusOK)
}
