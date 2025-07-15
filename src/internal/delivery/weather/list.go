package weather

import (
	"github.com/goriiin/rbs-test-task/src/internal/utils"
	"log"
	"net/http"
)

func (wd *WeatherDelivery) List(w http.ResponseWriter, r *http.Request) {
	weatherData, err := wd.repo.GetAll(r.Context())
	if err != nil {
		log.Printf("[ WeatherDelivery.List ] ERROR: failed to get weather data: %v", err)

		utils.WriteJSON(w,
			http.StatusInternalServerError,
			requestError{Error: err.Error()})

		return
	}

	w.Header().Set(key, value)

	err = wd.weather.Execute(w, weatherData)
	if err != nil {
		log.Printf("[ WeatherDelivery.List ] ERROR: failed to execute: %v", err)

		utils.WriteJSON(w, http.StatusInternalServerError, requestError{Error: internalServerError})
	}
}
