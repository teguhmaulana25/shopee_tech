package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguhmaulana25/shopee_tech/controller"
)

// NewRouter is function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
	var exchangeController = controller.NewExchangeController()

	r := mux.NewRouter().StrictSlash(true)
	r.Methods("GET").Path("/").HandlerFunc(exchangeController.Index)
	r.Methods("GET").Path("/exchange/all-data").HandlerFunc(exchangeController.All)
	r.Methods("POST").Path("/exchange/create").HandlerFunc(exchangeController.Store)
	r.Methods("DELETE").Path("/exchange/delete/{from}/{to}").HandlerFunc(exchangeController.Delete)
	r.Methods("GET").Path("/exchange/tracked/{date}").HandlerFunc(exchangeController.Tracked)
	r.Methods("GET").Path("/exchange/find/{from}/{to}").HandlerFunc(exchangeController.Find)
	http.Handle("/", r)
	return r
}
