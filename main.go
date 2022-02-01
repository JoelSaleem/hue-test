package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/amimof/huego"
	tm "github.com/buger/goterm"
)

// TODO ENV VARS
// TODO PATH TO FAVOURITES
const (
	USERNAME   = "PWljrRcSocWuN-2XL4bosQA5Pr9aNtmjuA6-cqlN"
	IP_ADDRESS = "192.168.68.106"
)

func main() {
	tm.Clear()

	bridge := huego.New(IP_ADDRESS, USERNAME)

	scenes, err := bridge.GetScenes()
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	// TODO: HANDLE CASE WITH NO FAVES
	faves := make(map[string]bool)
	for {
		tm.Clear()
		tm.Flush()
		tm.MoveCursor(1, 1)
		bytes, err := ioutil.ReadFile("faves")
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			} else {
				fmt.Println("no favourites, loading all")
			}
		}

		if len(bytes) > 0 {
			data := strings.Split(string(bytes), "\n")
			for i := range data {
				faves[data[i]] = true
				// faves = append(faves, data[i])
			}
		}

		s := []huego.Scene{}
		for _, scene := range scenes {
			if len(faves) == 0 || faves[scene.Name] {
				s = append(s, scene)
			}
		}

		listScenes(s)

		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		for _, s := range scenes {
			if s.Name == input {
				s.Recall(1)
			}
		}

	}

}
