package main

import (
	"fmt"

	"github.com/amimof/huego"
)

func listScenes(scenes []huego.Scene) {
	for _, s := range scenes {
		fmt.Println(s.Name)
	}
	fmt.Print("\n\n")
}
