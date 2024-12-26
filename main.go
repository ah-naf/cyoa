package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

type StoryOption struct {
	Text string `json: "text"`
	Arc  string `json: "arc"`
}

type Story struct {
	Title   string        `json:"title"`
	Story   []string      `json: "story"`
	Options []StoryOption `json: "options"`
}

type Adventure map[string]Story

var story Adventure

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func loadStory() {
	file, err := os.Open("story.json")
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&story)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v\n", err)
	}
}

func main() {
	loadStory()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := ""
		if r.URL.Path == "/" {
			path = "/intro"
		} else {
			path = r.URL.Path[1:]
			fmt.Println(path)
			if v, ok := story[path]; ok {
				tpl.Execute(w, v)
			} else {
				http.NotFound(w, r)
			}
		}
	})

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
