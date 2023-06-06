package api

import (
	"net/http"
)

type API struct {
	Router    *http.ServeMux
	navitagor Navigator
}

func Initialize(navigator Navigator) *API {

	api := API{
		Router:    http.NewServeMux(),
		navitagor: navigator,
	}

	api.Router.HandleFunc("/routes", api.handleRoutesRequest)

	return &api
}
