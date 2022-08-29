package clients

import (
	"fmt"
	"gettingLogs/internal/logger"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DataBase interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) error
}

type dataBase struct {
	db     *sqlx.DB
	logger logger.Logger
}

func New(logger logger.Logger) DataBase {

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWD")
	dbname := os.Getenv("MYSQL_DBNAME")

	source := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v",
		user, pass, host, port, dbname,
	)

	conn, err := sqlx.Connect("mysql", source)
	if err != nil {
		logger.Panic(err)
	}

	logger.Log("db connected")

	conn.SetConnMaxIdleTime(5 * time.Second)
	conn.SetConnMaxLifetime(60 * time.Second)
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(10)

	return &dataBase{
		db:     conn,
		logger: logger,
	}
}

func (d *dataBase) Select(dest interface{}, query string, args ...interface{}) error {
	return d.db.Select(dest, query, args...)
}

func (d *dataBase) Exec(query string, args ...interface{}) error {
	_, err := d.db.Exec(query, args...)
	return err
}
