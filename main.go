package main

import (
	"flag"
	"fmt"
	web "genshincal/web"
)

var addr = flag.String("addr", "localhost:8080", "server address and port")

func main() {
	fmt.Println("test")
	// genshindata.Download()
	// fmt.Printf("%#v", data.GetAvatarByName("安柏").LevelMap["90"])

	flag.Parse()
	web.Start(*addr)
}
