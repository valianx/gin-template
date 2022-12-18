package main

import (
	"github.com/valianx/gin-template/cmd/routes"
	"github.com/valianx/gin-template/pkg/config"
)

func main() {
	port, err := config.GetPort()
	if err != nil {
		return
	}

	r := routes.SetupRoutes(port)
	err = r.Run()
	if err != nil {
		return
	}

}
