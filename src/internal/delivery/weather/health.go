package weather

import (
	"github.com/goriiin/rbs-test-task/src/internal/utils"
	"log"
	"net/http"
)

func (wd *WeatherDelivery) Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(key, value)

	err := wd.healthy.Execute(w, nil)
	if err != nil {
		log.Println("[ WeatherDelivery.Health ]  ERROR: failed to execute", err)
		utils.WriteJSON(w, http.StatusInternalServerError, requestError{Error: internalServerError})
	}
}
