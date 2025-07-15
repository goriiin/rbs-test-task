package postgres_pool

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

func testPing(wg *sync.WaitGroup, pools map[string]*pgxpool.Pool) chan error {
	errChan := make(chan error, len(pools))

	for name, pool := range pools {
		wg.Add(1)

		go func(dbName string, dbPool *pgxpool.Pool) {
			defer wg.Done()

			if err := dbPool.Ping(context.Background()); err != nil {
				errChan <- fmt.Errorf("[ postgres_pool.testPing ] ERROR: ping failed for database '%s': %w", dbName, err)
			}
		}(name, pool)
	}

	wg.Wait()
	close(errChan)

	return errChan
}
