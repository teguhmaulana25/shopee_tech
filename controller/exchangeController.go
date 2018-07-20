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
		Data:    "Index",
	}

	renderJSON(w, response, http.StatusOK)
}

func (ExchangeController) All(w http.ResponseWriter, r *http.Request) {
	query := exchange.All()
	response := res{
		Code:    200,
		Message: "Success",
		Data:    query,
	}

	renderJSON(w, response, http.StatusOK)
}
