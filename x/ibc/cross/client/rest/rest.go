package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(ctx client.Context, r *mux.Router) {
	r.HandleFunc("/cross/coordinator/{tx_id}", QueryCoordinatorStatusHandlerFn(ctx)).Methods("GET")
	r.HandleFunc("/cross/unackpackets", QueryUnacknowledgedPacketsHandlerFn(ctx)).Methods("GET")
}
