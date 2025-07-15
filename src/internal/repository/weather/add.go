package weather

import (
	"context"
	"fmt"
	"github.com/goriiin/rbs-test-task/src/internal/domain"
)

func (r *WeatherRepository) Add(ctx context.Context, weather domain.Weather) error {
	query := `INSERT INTO weather (city, temperature) VALUES ($1, $2);`

	_, err := r.pool.Exec(ctx, query, weather.City, weather.Temperature)
	if err != nil {
		return fmt.Errorf("failed to execute insert for weather data: %w", err)
	}

	return nil
}
