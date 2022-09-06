package main

import (
	"pos-backend/routes"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	routes.Routes()
}
