package main

import (
	"github.com/tranvinh21/fastext-be-go/cmd/api"
	"github.com/tranvinh21/fastext-be-go/config"
	"github.com/tranvinh21/fastext-be-go/db"
)

func main() {

	db := db.NewDB()
	port := config.Envs.Server.PORT
	apiServer := api.NewAPIServer(db, port)
	apiServer.Run()
}
