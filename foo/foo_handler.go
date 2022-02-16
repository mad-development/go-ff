package foo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetFoo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := "Foo >> "
	for _, item := range params {
		data += fmt.Sprintf(" %s | ", item)
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error(err.Error())
		return

	}
	log.Info(data)
}
