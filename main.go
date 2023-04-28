package main

import (
	route "api-go/router"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Print("Server Loaded")
	route.StartServer().Run(":8086")
}
