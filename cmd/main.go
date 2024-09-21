package main

import (
	"fmt"

	"github.com/QuinsZouls/practica-6/cmd/config"
)

func main() {
	fmt.Println("Variables de entorno: ", config.USERNAME, config.PASSWORD)
}
