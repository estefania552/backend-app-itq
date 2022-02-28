package main

import (
	"log"

	"github.com/estefania552/backend-app-itq/bd"
	"github.com/estefania552/backend-app-itq/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
