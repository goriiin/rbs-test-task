package weather

import (
	"github.com/goriiin/rbs-test-task/src/internal/utils"
	"log"
	"net/http"
)

func (wd *WeatherDelivery) Show(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(key, value)

	err := wd.add.Execute(w, nil)
	if err != nil {
		log.Println("[ WeatherDelivery.Show ]  ERROR: failed to execute", err)
		utils.WriteJSON(w, http.StatusInternalServerError, requestError{Error: internalServerError})
	}
}
