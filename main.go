package main

import (
	"consumer/test/config"
	"consumer/test/pkg"
)

func main() {
	srv := config.New()
go pkg.Consume()
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
