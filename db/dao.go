/**
  @author: honor
  @since: 2021/7/7
  @desc: //TODO
**/
package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

type Dao struct {
	db *sqlx.DB
}

var DatabaseLink string

func (d Dao) connect() {
}
