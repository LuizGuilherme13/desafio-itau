package main

import (
	"flag"
	"log"

	"github.com/LuizGuilherme13/desafio-itau/api"
)

func main() {

	addr := flag.String("addr", ":8080", "server address")

	server := api.NewServer(*addr)

	log.Println(server.Start())

}
