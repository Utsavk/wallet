package main

import (
	"fmt"
	"wallet/cmd"
	"wallet/db/mysql"
	"wallet/server"
)

func main() {
	fmt.Printf("Application starts")
	var server = &server.Server{}
	cmd.EnvConf.ParseEnvVars()
	mysql.Conn.Connect()
	server.Configure()
	server.Start()
}
