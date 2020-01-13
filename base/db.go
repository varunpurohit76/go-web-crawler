package base

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func buildDSN() string {
	conf := Config.DB
	dataSourceName := conf.User + ":" + conf.Pwd + "@tcp(" + conf.Host + ":3306)/" + conf.Database + "?" + strings.Join(conf.Params, "&")
	log.Println("Connecting to DB using DSN: ", dataSourceName)
	return dataSourceName
}

func ConnectDb() (err error) {
	conf := Config.DB
	db, err = sqlx.Open("mysql", buildDSN())
	if err != nil {
		return fmt.Errorf("failed to Establish DB connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("db ping failed: %v", err)
	}
	db.SetMaxOpenConns(conf.MaxOpenConn)
	db.SetMaxIdleConns(conf.MaxIdleConn)
	return nil
}

func NewDbTransaction() (*sqlx.Tx, error) {
	return db.Beginx()
}
