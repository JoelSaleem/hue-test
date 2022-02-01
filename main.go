package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/amimof/huego"
	tm "github.com/buger/goterm"
)

const (
	USERNAME   = ""
	IP_ADDRESS = ""
)

func main() {
	tm.Clear()

	bridge := huego.New(IP_ADDRESS, USERNAME)

	scenes, err := bridge.GetScenes()
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		for _, s := range scenes {
			fmt.Println(s.Name)
		}
		fmt.Print("\n\n")

		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		for _, s := range scenes {
			if s.Name == input {
				s.Recall(1)
			}
		}
		tm.Clear()
		tm.Flush()
	}

}
