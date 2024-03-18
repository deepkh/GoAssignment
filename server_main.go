package main

import (
	"go-recommendation-system/api"
	"go-recommendation-system/services"
)

func main() {
	go api.RunApiServer()
	services.RunGrpcServer()
}
