package main

import (
	"flag"
	"fmt"

	// data "genshincal/genshindata"
	web "genshincal/web"
)

var addr = flag.String("addr", "localhost:8080", "server address and port")

func main() {
	fmt.Println("test")
	// genshindata.Download()
	// fmt.Printf("%#v\n", data.GetAvatarByName("安柏").LevelMap["90"])
	// fmt.Printf("%#v\n", data.GetWeaponByName("天空之刃").LevelMap["90"])

	flag.Parse()
	web.Start(*addr)
}
