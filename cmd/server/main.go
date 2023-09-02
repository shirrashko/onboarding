package main

import (
	"github.com/pkg/errors"
	"github.com/shirrashko/BuildingAServer-step2/config"
	"github.com/shirrashko/BuildingAServer-step2/pkg/api"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(errors.Wrap(err, "error loading config"))
	}

	server, err := api.NewServer(conf, api.Router)
	if err != nil {
		panic(errors.Wrap(err, "Unable to create server"))
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(errors.Wrap(err, "Unable to start the server"))
	}
}
