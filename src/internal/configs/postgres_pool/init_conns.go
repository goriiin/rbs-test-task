package postgres_pool

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

func InitConnections(configs []DSNConfig) (map[string]*pgxpool.Pool, error) {
	pools := make(map[string]*pgxpool.Pool)
	for _, cfg := range configs {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

		poolCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			return nil, fmt.Errorf(
				"[ postgres_pool.InitConnections ] ERROR: failed to parse DSN for '%s': %w",
				cfg.Name, err)
		}

		poolCfg.MaxConns = cfg.MaxConns
		poolCfg.MinConns = cfg.MinConns
		poolCfg.MaxConnLifetime = cfg.MaxConnLifetime
		poolCfg.MaxConnIdleTime = cfg.MaxConnIdleTime

		pool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
		if err != nil {
			return nil, fmt.Errorf(
				"[ postgres_pool.InitConnections ] ERROR: failed to create connection pool for '%s': %w",
				cfg.Name, err)
		}

		pools[cfg.Name] = pool
	}

	errChan := make(chan error, len(pools))
	var wg sync.WaitGroup
	wg.Add(len(pools))

	for name, pool := range pools {
		go func(dbName string, dbPool *pgxpool.Pool) {
			defer wg.Done()
			if err := dbPool.Ping(context.Background()); err != nil {
				errChan <- fmt.Errorf(
					"[ postgres_pool.InitConnections ] ERROR: ping failed for required database '%s': %w",
					dbName, err)
			}
		}(name, pool)
	}

	wg.Wait()
	close(errChan)

	var allErrors []error
	for err := range errChan {
		allErrors = append(allErrors, err)
	}

	if len(allErrors) > 0 {
		return nil, errors.Join(allErrors...)
	}

	return pools, nil
}
