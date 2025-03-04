package main

import (
	"flag"

	"github.com/LuizGuilherme13/desafio-itau/internal/server"
)

func main() {

	addr := flag.String("addr", ":8080", "server address")

	server.New(*addr).Start()
}
