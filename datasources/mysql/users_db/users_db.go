package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_user         = "mysql_user"
	mysql_password     = "mysql_password"
	mysql_host         = "mysql_host"
	mysql_port         = "mysql_port"
	mysql_users_scheme = "mysql_users_scheme"
)

var (
	Client *sql.DB

	user     = os.Getenv(mysql_user)
	password = os.Getenv(mysql_password)
	host     = os.Getenv(mysql_host)
	port     = os.Getenv(mysql_port)
	scheme   = os.Getenv(mysql_users_scheme)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		user, password, host, port, scheme,
	)

	log.Printf("about to connect to %s", dataSourceName)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
