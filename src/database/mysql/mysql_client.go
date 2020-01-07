package mysql

import (
	"database/sql"
	"fmt"
	"github.com/Teslenk0/utils-go/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysqlInventoryUsername = "mysql_inventory_username"
	mysqlInventoryPassword = "mysql_inventory_password"
	mysqlInventoryHost     = "mysql_inventory_host"
	mysqlInventorySchema   = "mysql_inventory_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlInventoryUsername)
	password = os.Getenv(mysqlInventoryPassword)
	host     = os.Getenv(mysqlInventoryHost)
	schema   = os.Getenv(mysqlInventorySchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	//mysql.SetLogger(logger.GetLogger())
	logger.Info("database succesfuly configured")
}
