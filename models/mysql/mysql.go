package mysql

import (
	"github.com/alimy/chi-music/models/core"
	"github.com/alimy/chi-music/models/model"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // mysql sql driver
)

type mysqlRepository struct {
	*core.Sqlx
}

// NewRepository build a new core.Repository that backend by mysql database
func NewRepository(config *model.Config) (core.Repository, error) {
	_, dsn := config.Dsn()
	mysqlDb, err := sqlx.Connect("mysql", dsn)
	return &mysqlRepository{Sqlx: &core.Sqlx{DB: mysqlDb}}, err
}
