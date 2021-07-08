/**
  @author: honor
  @since: 2021/7/7
  @desc: //TODO
**/
package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

var dao *Dao

type Dao struct {
	db *sqlx.DB
}

const (
	ddl = `
	CREATE TABLE IF NOT EXISTS  bill(
			id                                      BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键编号',
			channel									BIGINT         NOT NULL,
			paid_time								TIMESTAMP      DEFAULT NULL,
			charge                                  VARCHAR(64)    NOT NULL DEFAULT '' COMMENT '账单总金额',
			status								    INTEGER        DEFAULT NULL,
			version                                 INT            DEFAULT 0,
			removed                                 BIT(1)         DEFAULT 0
 )`
	dml = ``
)

func (dao *Dao) connect(dbLink string) (err error) {
	if dbLink == "" {
		panic("DatabaseLink not set")
	}
	if strings.Index(dbLink, "@") > 0 {
		var d string
		if strings.Index(dbLink, "?") > 0 {
			d = dbLink + "parseTime=true"
		} else {
			d = dbLink + "?parseTime=true"
		}
		var er error
		dao.db, er = sqlx.Open("mysql", d)
		if er == nil {
			_, er = dao.db.Exec(ddl)
			if er != nil {
				log.Printf("Init db failed. %s", er)
			}
		}
		dao.db.SetMaxIdleConns(5)
		dao.db.SetMaxOpenConns(10)
	}
	panic("unknown datan base link")
}

func (dao *Dao) Create(query string, args ...interface{}) (id uint64, er error) {
	if r, err := dao.db.Exec(query, args...); err != nil {
		return 0, err
	} else {
		var insertId int64
		insertId, er = r.LastInsertId()
		return uint64(insertId), er
	}
}

func (dao *Dao) Update(query string, args ...interface{}) (rows uint64, er error) {
	if r, err := dao.db.Exec(query, args...); err != nil {
		return 0, er
	} else {
		var rowsAffected int64
		rowsAffected, er = r.RowsAffected()
		return uint64(rowsAffected), er
	}
}

func (dao *Dao) LoadById(query string, id uint64, mapping interface{}) (err error) {
	return dao.db.Select(mapping, query, id)
}
