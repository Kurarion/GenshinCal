package main

import (
	"flag"
	"fmt"
	web "genshincal/web"
)

var addr = flag.String("addr", "localhost:8080", "example [address:port]")

func main() {
	fmt.Println("Started!")
	flag.Parse()
	web.Start(*addr)
	select {}
}
