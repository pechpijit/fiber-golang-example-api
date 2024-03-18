package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pechpijit/Fiber_golang_example_api/database/query"
	"github.com/pechpijit/Fiber_golang_example_api/utils"
)

type QueriesProduct struct {
	*query.ProductQuery // load queries from Book model
}

func OpenDBConnection() (*QueriesProduct, error) {
	connectString, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error, can not ping to database, %w", err)
	}

	return &QueriesProduct{
		ProductQuery: &query.ProductQuery{DB: db}, // from Book model
	}, nil
}
