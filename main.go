package main

import (
	"fmt"

	"github.com/chagasduarte/go-rest-api/database"
	"github.com/chagasduarte/go-rest-api/routes"
)

func main() {
	database.ConectBase()
	fmt.Println("Servidor conectado")
	routes.HandleRequest()
}
