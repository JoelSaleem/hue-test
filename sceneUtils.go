package main

import (
	"log"

	"github.com/amimof/huego"
)

func getScenesFromFavourites(scenes []huego.Scene, favourites Favourites) Scenes {
	s := []huego.Scene{}

	for _, scene := range scenes {
		if len(favourites) == 0 || favourites[scene.Name] {
			if favourites[scene.Name] {
				favourites[scene.Name] = false
			}
			s = append(s, scene)
		}
	}

	return s
}

func turnAllLightsOff(b huego.Bridge) {
	lights, err := b.GetLights()
	if err != nil {
		log.Fatalln(err)
	}

	for _, l := range lights {
		l.Off()
	}
}
