package bar

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetBar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := "Bar >> "
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
