package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var Stories Story

func main() {

	file, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	s := json.NewDecoder(file)

	s.Decode(&Stories)

	//http.HandleFunc("/", handler)
	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.Execute(w, Stories["intro"])
	// })
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// for k := range stories {
// 	fmt.Fprintf(w, stories[k].Title)
// }
func test(w http.ResponseWriter, r *http.Request) {
	s := strings.TrimPrefix(r.RequestURI, "/")
	fmt.Println(s)

	if s == "" {
		s = "intro"
	}

	tmpl := template.Must(template.ParseFiles("./pages/index.gtpl"))
	tmpl.Execute(w, Stories[s])
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Stories["intro"].Title)
	fmt.Fprintln(w)
	fmt.Fprintln(w)

	for i := range Stories["intro"].Story {
		fmt.Fprintf(w, Stories["intro"].Story[i])
		fmt.Fprintln(w)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "Options:")
	for i := range Stories["intro"].Options {
		fmt.Fprintf(w, Stories["intro"].Options[i].Text)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintf(w, "Button Nr %v:", i)
		fmt.Fprintln(w)
		fmt.Fprintf(w, Stories["intro"].Options[i].Arc)

		fmt.Fprintln(w)
		fmt.Fprintln(w)
	}

}
