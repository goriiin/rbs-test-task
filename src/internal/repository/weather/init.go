package weather

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type WeatherRepository struct {
	pool *pgxpool.Pool
}

func NewWeatherRepository(pool *pgxpool.Pool) *WeatherRepository {
	return &WeatherRepository{pool: pool}
}
