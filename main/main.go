package main

import (
	db "github.com/Yq2/gin_demo_api/database"
	route_handler "github.com/Yq2/gin_demo_api/router"
)

func main() {
	defer db.SqlDB.Close()
	router := route_handler.InitRouter()
	router.Run(":8989")
}
