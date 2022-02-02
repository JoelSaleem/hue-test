package main

import (
	"fmt"

	"github.com/amimof/huego"
)

type Scenes []huego.Scene

func (scenes Scenes) ListScenes() {
	for _, s := range scenes {
		fmt.Println(s.Name)
	}
	fmt.Print("\n\n")
}

func (scenes Scenes) ActivateScene(input string) {
	for _, s := range scenes {
		if s.Name == input {
			s.Recall(1)
		}
	}
}
