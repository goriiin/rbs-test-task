package postgres_pool

import (
	"fmt"
	"os"
	"time"
)

const (
	defaultMaxConnections        = 10
	defaultMinConnections        = 0
	defaultMaxConnectionLifeTime = time.Hour * 2
	defaultMinConnectionIdleTime = time.Minute * 30
)

func buildFinalConfig(template DBConfigTemplate) (*DSNConfig, error) {
	maxConns := int32(defaultMaxConnections)
	minConns := int32(defaultMinConnections)
	maxConnLifetime := defaultMaxConnectionLifeTime
	minConnIdleTime := defaultMinConnectionIdleTime

	if template.PoolSettings != nil {
		if template.PoolSettings.MaxConnections != nil {
			maxConns = *template.PoolSettings.MaxConnections
		}

		if template.PoolSettings.MinConnections != nil {
			minConns = *template.PoolSettings.MinConnections
		}

		if template.PoolSettings.MaxConnectionLifetime != nil {
			parsed, err := time.ParseDuration(*template.PoolSettings.MaxConnectionLifetime)
			if err != nil {
				return nil, fmt.Errorf("[ postgres_pool.buildFinalConfig ] ERROR: invalid max_connection_lifetime for %s: %w",
					template.Name, err)
			}

			maxConnLifetime = parsed
		}
		if template.PoolSettings.MinConnectionIdleTime != nil {
			parsed, err := time.ParseDuration(*template.PoolSettings.MinConnectionIdleTime)
			if err != nil {
				return nil, fmt.Errorf("[ postgres_pool.buildFinalConfig ] ERROR: invalid min_connection_idle_time for %s: %w",
					template.Name, err)
			}

			minConnIdleTime = parsed
		}
	}

	password := os.Getenv(template.EnvVars.Password)
	if password == "" {
		return nil, fmt.Errorf("[ postgres_pool.buildFinalConfig ] ERROR: environment variable for password %s is not set",
			template.EnvVars.Password)
	}

	return &DSNConfig{
		Name:            template.Name,
		Host:            template.Host,
		Port:            template.Port,
		User:            os.Getenv(template.EnvVars.User),
		Password:        password,
		DBName:          os.Getenv(template.EnvVars.DBName),
		MaxConns:        maxConns,
		MinConns:        minConns,
		MaxConnLifetime: maxConnLifetime,
		MaxConnIdleTime: minConnIdleTime,
	}, nil
}
