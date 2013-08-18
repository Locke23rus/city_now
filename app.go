package main

import (
	"encoding/json"
	eventsource "github.com/antage/eventsource/http"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	publicPath       = "public"
	rootTemplatePath = "templates/index.html"
	eventID          = ""
	statsEvent       = "stats"
)

var (
	rootTemplate *template.Template
	feed         eventsource.EventSource
	stats        = Stats{0}
)

type Stats struct {
	Online int `json:"online"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, nil)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, publicPath+r.URL.Path)
}

func sendStats() {
	stats.Online = feed.ConsumersCount()
	json, err := json.Marshal(stats)
	if err != nil {
		log.Printf("Can't marshal Stats: %v", err)
		return
	}
	feed.SendMessage(string(json), statsEvent, eventID)
}

func statsPublisher() {
	for {
		sendStats()
		time.Sleep(2 * time.Second)
	}
}

func init() {
	rootTemplate = template.Must(template.ParseFiles(rootTemplatePath))
}

func main() {
	http.HandleFunc("/css/", staticHandler)
	http.HandleFunc("/js/", staticHandler)
	http.HandleFunc("/", rootHandler)

	feed = eventsource.New(nil)
	defer feed.Close()
	http.Handle("/feed", feed)
	go statsPublisher()

	log.Fatal(http.ListenAndServe(":3000", nil))
}
