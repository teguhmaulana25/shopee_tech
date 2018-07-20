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
	http.Handle("/", r)
	return r
}
