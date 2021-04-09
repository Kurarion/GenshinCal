package main

import (
	"fmt"
	data "genshincal/genshindata"
)

func main() {
	fmt.Println("test")
	// genshindata.Download()
	fmt.Printf("%#v", data.GetAvatarByName("安柏").LevelMap["90"])
}
