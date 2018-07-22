package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguhmaulana25/shopee_tech/config/middleware"
	"github.com/teguhmaulana25/shopee_tech/controller"
)

// NewRouter is function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
	var exchangeController = controller.NewExchangeController()
	var userController = controller.NewUserController()

	r := mux.NewRouter().StrictSlash(true)
	// Log server
	r.Use(middleware.Logger)
	r.Use(middleware.Security)

	// Auth
	r.Methods("POST").Path("/auth").HandlerFunc(userController.Auth)

	// Exchange
	r.Methods("GET").Path("/").HandlerFunc(exchangeController.Index)
	r.Methods("GET").Path("/exchange/all-data").HandlerFunc(middleware.JwtAuth(exchangeController.All))
	r.Methods("POST").Path("/exchange/create").HandlerFunc(middleware.JwtAuth(exchangeController.Store))
	r.Methods("DELETE").Path("/exchange/delete/{from}/{to}").HandlerFunc(middleware.JwtAuth(exchangeController.Delete))
	r.Methods("GET").Path("/exchange/tracked/{date}").HandlerFunc(middleware.JwtAuth(exchangeController.Tracked))
	r.Methods("GET").Path("/exchange/find/{from}/{to}").HandlerFunc(middleware.JwtAuth(exchangeController.Find))
	http.Handle("/", r)
	return r
}
