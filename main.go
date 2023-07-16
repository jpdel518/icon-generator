package main

import (
	"github.com/jpdel518/icon_generator/handler"
	"github.com/jpdel518/icon_generator/service"
)

func main() {
	iconService := service.NewIconService()
	identiconSerivce := service.NewIdenticonService()
	handler.NewHandler(iconService, identiconSerivce)
}
