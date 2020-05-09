package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/roneetkumar/cyoa"
)

func main() {
	fmt.Println("Adventurous Books")

	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")

	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)

	var story cyoa.Story

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
