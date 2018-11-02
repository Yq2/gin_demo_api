package main

import (
	db "gin_demo_api/database"
	route_handler "gin_demo_api/router"
)

func main() {
	defer db.SqlDB.Close()
	router := route_handler.InitRouter()
	router.Run(":8989")
}
