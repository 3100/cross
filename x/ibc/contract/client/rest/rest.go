package rest

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(ctx client.Context, r *mux.Router) {
	r.HandleFunc("/cross/contract/call", QueryContractCallRequestHandlerFn(ctx)).Methods("POST")
}
