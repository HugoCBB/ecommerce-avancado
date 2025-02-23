package main

import (
	"github.com/HugoCBB/Api/database"
	"github.com/HugoCBB/Api/routers"
)

func main() {
	database.ConnectDatabase()
	routers.HandleRequest()
}
