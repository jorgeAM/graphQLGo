package db

import (
	"sync"

	"github.com/go-pg/pg/v10"
)

var (
	dbHandler *PostgresHandler
	once      sync.Once
	someError error
)

// PostgresHandler handle operation tu postgres database
type PostgresHandler struct {
	DB *pg.DB
}

// NewPostgresHandler initialize postgres handler
func NewPostgresHandler(url string) (Handler, error) {
	once.Do(func() {
		opt, err := pg.ParseURL(url)
		someError = err
		db := pg.Connect(opt)
		dbHandler = &PostgresHandler{DB: db}
	})

	return dbHandler, someError
}

// Close method closes postgres connection
func (h *PostgresHandler) Close() error {
	return h.DB.Close()
}
