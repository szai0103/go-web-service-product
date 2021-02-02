package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/schwarz/inventoryservice/product/api"
	"github.com/schwarz/inventoryservice/product/data/mysql"
	api2 "github.com/schwarz/inventoryservice/receipt/api"
	"net/http"
)

const apiBasePath = "/api"

func main() {
	mysql.SetupDB()
	api.SetupRoutes(apiBasePath)
	api2.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
