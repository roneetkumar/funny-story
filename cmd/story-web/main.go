package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	funnystory "github.com/roneetkumar/funny-story"
)

func main() {
	fmt.Println("Adventurous Books")

	port := flag.Int("port", 3000, "the port to start the server")
	filename := flag.String("file", "gopher.json", "the JSON file with the funny story")

	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := funnystory.JSONStory(f)

	if err != nil {
		panic(err)
	}

	h := funnystory.NewHandler(story)

	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

	fmt.Printf("%+v\n", story)

}
