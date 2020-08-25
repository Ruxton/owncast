package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gabek/owncast/config"
	"github.com/gabek/owncast/router/middleware"
)

//GetWebConfig gets the status of the server
func GetWebConfig(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)

	configuration := config.Config.InstanceDetails
	configuration.Version = config.Config.VersionInfo
	json.NewEncoder(w).Encode(configuration)
}
