package osrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LuddesGit/there-yet/internal/api"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) GetRoute(source, destination *api.Location) (api.Route, error) {
	url := fmt.Sprintf("%s/%f,%f;%f,%f?overview=false", c.baseURL, source.Lng, source.Lat, destination.Lng, destination.Lat)

	resp, err := http.Get(url)
	if err != nil {
		return api.Route{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return api.Route{}, fmt.Errorf("failed to fetch routes: %s", resp.Status)
	}

	var osrmResp struct {
		Routes []struct {
			Duration float64 `json:"duration"`
			Distance float64 `json:"distance"`
		} `json:"routes"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	err = json.NewDecoder(resp.Body).Decode(&osrmResp)
	if err != nil {
		return api.Route{}, fmt.Errorf("failed to decode OSRM response: %w", err)
	}

	if osrmResp.Code != "Ok" || len(osrmResp.Routes) == 0 {
		return api.Route{}, fmt.Errorf("failed to fetch from OSRM: %s", osrmResp.Message)
	}

	route := api.Route{
		Destination: fmt.Sprintf("%f,%f", destination.Lat, destination.Lng),
		Duration:    osrmResp.Routes[0].Duration,
		Distance:    osrmResp.Routes[0].Distance,
	}

	return route, nil
}
