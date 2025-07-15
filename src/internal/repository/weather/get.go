package weather

import (
	"context"
	"fmt"
	"github.com/goriiin/rbs-test-task/src/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (r *WeatherRepository) GetAll(ctx context.Context) ([]domain.Weather, error) {
	query :=
		`SELECT city, temperature 
		 FROM weather 
		 ORDER BY city;`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[ WeatherRepository.GetAll ] ERROR: failed to execute query for weather data: %w", err)
	}
	defer rows.Close()

	weatherData := make([]domain.Weather, 0, 10)

	_, err = pgx.ForEachRow(rows, []any{new(string), new(int)}, func() error {
		var w domain.Weather
		if err = rows.Scan(&w.City, &w.Temperature); err != nil {
			return err
		}

		weatherData = append(weatherData, w)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("[ WeatherRepository.GetAll ] ERROR: failed to scan weather data rows: %w", err)
	}

	return weatherData, nil
}
