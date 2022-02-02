package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Favourites map[string]bool

func (f Favourites) ReadFaves() Favourites {
	bytes, err := ioutil.ReadFile(FAVOURITES_PATH)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}

	}

	if len(bytes) > 0 {
		data := strings.Split(string(bytes), "\n")
		for i := range data {
			f[data[i]] = true
		}
	} else {
		fmt.Print("no favourites, loading all\n\n")
	}

	return f
}
