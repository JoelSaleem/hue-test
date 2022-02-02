package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/amimof/huego"
	tm "github.com/buger/goterm"
)

// TODO ENV VARS
// TODO PATH TO FAVOURITES
const (
	USERNAME        = "PWljrRcSocWuN-2XL4bosQA5Pr9aNtmjuA6-cqlN"
	IP_ADDRESS      = "192.168.68.106"
	FAVOURITES_PATH = "/Users/joelsaleem/.hue-test/faves"
)

func main() {
	tm.Clear()

	bridge := huego.New(IP_ADDRESS, USERNAME)

	scenes, err := bridge.GetScenes()
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	favourites := Favourites{}
	for {
		tm.Clear()
		tm.Flush()
		tm.MoveCursor(1, 1)

		favourites = favourites.ReadFaves()

		s := getScenesFromFavourites(scenes, favourites)
		s.ListScenes()

		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if input == "Off" {
			turnAllLightsOff(*bridge)
		} else {
			s.ActivateScene(input)
		}

	}

}
