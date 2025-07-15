package weather

import (
	"github.com/goriiin/rbs-test-task/src/internal/utils"
	"net/http"
)

func (wd *WeatherDelivery) Ping(w http.ResponseWriter, _ *http.Request) {
	utils.WriteJSON(w, http.StatusOK, request{Status: "PONG"})
}
