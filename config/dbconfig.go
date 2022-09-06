package config

import "fmt"

/*var server = "LAPTOP-1V4QVAFB\\SQLEXPRESS"
var port = 1433
var user = "sa"
var password = "123456"
var database = "frank_ac"*/

var server = "202.133.4.238"
var port = 1433
var user = "sa"
var password = "Mlisql123"
var database = "frank_ac"

type DbConfig struct {
	ServerName   string
	PortNumber   int
	UserName     string
	Password     string
	DatabaseName string
}

func ConnDb() DbConfig {
	var f DbConfig
	f.ServerName = server
	f.PortNumber = port
	f.UserName = user
	f.Password = password
	f.DatabaseName = database
	return f
}

func ConnString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=disable;", server, user, password, port, database)
}
