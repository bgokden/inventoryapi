package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bgokden/inventoryapi/server"
)

func main() {
	var port = flag.Int("port", 8080, "Port for HTTP server")
	flag.Parse()
	e, err := server.CreateEchoServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// And we serve HTTP unless there is an error
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
