package api

type Navigator interface {
	GetRoute(source, destination *Location) (Route, error)
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Route struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}
