package weather

import (
	"github.com/goriiin/rbs-test-task/src/internal/domain"
	"github.com/goriiin/rbs-test-task/src/internal/utils"
	"net/http"
	"strconv"
)

func (wd *WeatherDelivery) Add(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, requestError{Error: badRequest})

		return
	}

	city := r.PostForm.Get("city")
	tempStr := r.PostForm.Get("temperature")

	if city == "" || tempStr == "" {
		utils.WriteJSON(w, http.StatusBadRequest, requestError{Error: required})

		return
	}

	temperature, err := strconv.Atoi(tempStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, requestError{Error: invalid})

		return
	}

	newWeather := domain.Weather{
		City:        city,
		Temperature: temperature,
	}

	if err = wd.repo.Add(r.Context(), newWeather); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, requestError{Error: internalServerError})

		return
	}

	http.Redirect(w, r, "/list", http.StatusSeeOther)
}
