package main

import (
	"log"
	"net/http"

	"github.com/LuddesGit/there-yet/internal/api"
	"github.com/LuddesGit/there-yet/internal/config"
	"github.com/LuddesGit/there-yet/internal/osrm"
)

func main() {
	cfg := config.LoadConfig()
	navigator := osrm.NewClient(cfg.OsrmURL)
	api := api.Initialize(navigator)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, api.Router))
}
