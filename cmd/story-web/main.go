package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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

	tmp := template.Must(template.New("").Parse(storyHandlerTeml))

	h := funnystory.NewHandler(
		story,
		funnystory.WithTemplate(tmp),
		funnystory.WithPathFunc(pathFn),
	)

	mux := http.NewServeMux()

	mux.Handle("/story/", h)

	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

	fmt.Printf("%+v\n", story)

}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyHandlerTeml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Funny Story</title>
</head>
<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
<body>
	 <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
</body>
</html>
`
