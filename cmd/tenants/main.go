package main

import (
	"flag"
	"fmt"

	"github.com/TylerGrey/tenants/internal/server"
)

var addr *string

func init() {
	addr = flag.String("http", ":8080", "HTTP server port")
	flag.Parse()
}

func main() {
	s := server.Server{
		Addr: addr,
	}

	if err := s.Start(); err != nil {
		fmt.Println(err.Error())
	}
}
