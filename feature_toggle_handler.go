package main

import (
	"github.com/Unleash/unleash-client-go/v3"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	token := "default:default.24daadb2e3182f04fbb26e66eed6fb01ee7ee26cb43df9702ca032a5"

	err := unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("default"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {token}}),
	)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Success getting toggle")
}

func IsEnabled(data string) bool {
	return unleash.IsEnabled(data)
}
