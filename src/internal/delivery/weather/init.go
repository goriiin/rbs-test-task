package weather

import (
	"context"
	"fmt"
	"github.com/goriiin/rbs-test-task/src/internal/domain"
	"html/template"
)

type requestError struct {
	Error string `json:"error"`
}

type request struct {
	Status string `json:"status"`
}

const (
	key                 = "Content-Type"
	value               = "text/html; charset=utf-8"
	internalServerError = "Internal server error"
	badRequest          = "Bad Request"
	required            = "City and temperature are required"
	invalid             = "Invalid temperature value"
)

type repo interface {
	GetAll(ctx context.Context) ([]domain.Weather, error)
	Add(ctx context.Context, weather domain.Weather) error
}

type WeatherDelivery struct {
	repo    repo
	healthy *template.Template
	weather *template.Template
	add     *template.Template
}

func NewWeatherDelivery(repo repo, templatePath string, healthyPath string, addPath string) (*WeatherDelivery, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("[ NewWeatherDelivery ] failed to parse %s: %w", templatePath, err)
	}

	tmplH, err := template.ParseFiles(healthyPath)
	if err != nil {
		return nil, fmt.Errorf("[ NewWeatherDelivery ] failed to parse  %s: %w", templatePath, err)
	}

	tmplA, err := template.ParseFiles(addPath)
	if err != nil {
		return nil, fmt.Errorf("[ NewWeatherDelivery ] failed to parse  %s: %w", templatePath, err)
	}

	return &WeatherDelivery{
		repo:    repo,
		weather: tmpl,
		healthy: tmplH,
		add:     tmplA,
	}, nil
}
