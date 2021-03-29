package main

import (
	"goTaxi/bootstrap"
	"goTaxi/driver/controller"
)

func main() {
	bootstrap.Start()
	controller.Index()
}
