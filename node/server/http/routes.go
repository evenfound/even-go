package http

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type successResponse struct {
	Data interface{} `json:"data"`
}

type appHandler func(w http.ResponseWriter, r *http.Request) *ErrorResponse

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s",
			e.Code, e.Message)

		http.Error(w, e.Message, e.Code)
	} else {
		fmt.Printf(
			"%s %s - %s \n",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

func NewRouter() *mux.Router {

	routers := mux.NewRouter().StrictSlash(true)

	routers.Methods("POST").Path("/newwallet").Handler(appHandler(NewWallet))
	routers.Methods("POST").Path("/newaccount").Handler(appHandler(NewAccount))
	routers.Methods("GET").Path("/listaccounts").Handler(appHandler(ListAccounts))
	routers.Methods("POST").Path("/newaddress").Handler(appHandler(NewWallet))

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, routers))

	return routers

}
