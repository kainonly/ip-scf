// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"ip2region-mongo/app"
	"ip2region-mongo/app/index"
	"ip2region-mongo/bootstrap"
	"ip2region-mongo/common"
)

// Injectors from wire.go:

func App(value *common.Values) (*gin.Engine, error) {
	client, err := bootstrap.UseMongoDB(value)
	if err != nil {
		return nil, err
	}
	database := bootstrap.UseDatabase(client, value)
	inject := &common.Inject{
		Values:      value,
		MongoClient: client,
		Db:          database,
	}
	service := &index.Service{
		Inject: inject,
	}
	controller := &index.Controller{
		Service: service,
	}
	engine := app.New(value, controller)
	return engine, nil
}
