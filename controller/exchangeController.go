package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/teguhmaulana25/shopee_tech/model"
)

var exchange = model.NewExchangeModel()

type ExchangeController struct{}

func NewExchangeController() ExchangeController {
	return ExchangeController{}
}

type filter struct {
	CurrentDate string
	LastDate    string
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

func (ExchangeController) Store(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Error(err)
		response := res{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}

		renderJSON(w, response, http.StatusUnprocessableEntity)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Error(err)
		response := res{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}

		renderJSON(w, response, http.StatusUnprocessableEntity)
		return
	}
	// Convert the JSON in the request to a Go type.
	if err := json.Unmarshal(body, &exchange); err != nil {
		log.Error(err)
		response := res{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}

		renderJSON(w, response, http.StatusUnprocessableEntity)
		return
	}
	exchange.CreatedByIp = GetIP()
	exchange.UpdatedByIp = GetIP()
	exchange.CreatedAt = string(time.Now().Format("2006-01-02 15:04:05"))
	exchange.UpdatedAt = string(time.Now().Format("2006-01-02 15:04:05"))

	query, err := exchange.Store()
	if err != nil {
		log.Debug(err)
		response := res{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}

		renderJSON(w, response, http.StatusUnprocessableEntity)
		return
	}
	response := res{
		Code:    200,
		Message: "Success",
		Data:    query,
	}

	renderJSON(w, response, http.StatusOK)
}

func (ExchangeController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Query the database.
	query := exchange.Delete(vars["from"], vars["to"])

	response := res{
		Code:    200,
		Message: "Success",
		Data:    query,
	}

	renderJSON(w, response, http.StatusOK)
}

func (ExchangeController) Tracked(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentDate := vars["date"]
	currentDateConvert, err := time.Parse("2006-01-02", currentDate)
	Check(err)

	getLastDate := currentDateConvert.AddDate(0, 0, -7)
	lastDate := getLastDate.Format("2006-01-02")

	// Query the database.
	query := exchange.Tracked(currentDate, lastDate)
	response := res{
		Code:    200,
		Message: "Success",
		Data:    query,
	}

	renderJSON(w, response, http.StatusOK)
}

func (ExchangeController) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currencyFrom := vars["from"]
	currencyTo := vars["to"]

	// Query the database.
	query := exchange.Find(currencyFrom, currencyTo)
	response := res{
		Code:    200,
		Message: "Success",
		Data:    query,
	}

	renderJSON(w, response, http.StatusOK)
}
