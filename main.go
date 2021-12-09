package main

import (
	"rizadwi.com/config"
	"rizadwi.com/routes"
)

func main() {
	// RunExperiment()
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
