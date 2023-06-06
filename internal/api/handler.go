package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type RoutesResponse struct {
	Source Location `json:"source"`
	Routes []Route  `json:"routes"`
}

func (a *API) handleRoutesRequest(w http.ResponseWriter, r *http.Request) {
	source, destinations, err := parseRequestURL(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	routes, err := a.getRoutes(source, destinations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sortRoutesByDuration(routes)

	response := RoutesResponse{
		Source: source,
		Routes: routes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func parseRequestURL(url *url.URL) (Location, []Location, error) {
	srcParam := url.Query().Get("src")
	dstParams := url.Query()["dst"]

	source, err := parseLocationParam(srcParam)
	if err != nil {
		return Location{}, nil, fmt.Errorf("invalid source location: %w", err)
	}

	var destinations []Location
	for _, dstParam := range dstParams {
		dst, err := parseLocationParam(dstParam)
		if err != nil {
			return Location{}, nil, fmt.Errorf("invalid destination location: %w", err)
		}
		destinations = append(destinations, dst)
	}

	return source, destinations, nil
}

func parseLocationParam(param string) (Location, error) {
	coords := strings.Split(param, ",")
	if len(coords) != 2 {
		return Location{}, fmt.Errorf("invalid location format")
	}

	lat, err := strconv.ParseFloat(coords[0], 64)
	if err != nil {
		return Location{}, fmt.Errorf("invalid latitude: %w", err)
	}

	lng, err := strconv.ParseFloat(coords[1], 64)
	if err != nil {
		return Location{}, fmt.Errorf("invalid longitude: %w", err)
	}

	return Location{Lat: lat, Lng: lng}, nil
}

func (a *API) getRoutes(source Location, destinations []Location) ([]Route, error) {
	var routes []Route

	for _, destination := range destinations {
		route, err := a.navitagor.GetRoute(&source, &destination)
		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	return routes, nil
}

func sortRoutesByDuration(routes []Route) {
	sort.Slice(routes, func(i, j int) bool {
		if routes[i].Duration == routes[j].Duration {
			return routes[i].Distance < routes[j].Distance
		}
		return routes[i].Duration < routes[j].Duration
	})
}
