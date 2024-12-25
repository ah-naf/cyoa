package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
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
	fmt.Println(story)
}
