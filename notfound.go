package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Saying struct {
	Sentence string `json:"sentence"`
	Author   string `json:"author"`
}

func loadSaying() (sayings []Saying, err error) {
	file, err := ioutil.ReadFile("./static/saying.json")
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(file, &sayings)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func notFoundHandler() http.HandlerFunc {

	var (
		once    sync.Once
		sayings []Saying
		t       *template.Template
	)

	return func(w http.ResponseWriter, r *http.Request) {

		once.Do(func() {

			var err error
			sayings, err = loadSaying()
			if err != nil {

				sayings = []Saying{
					{
						Sentence: "Catch my dream",
						Author:   "모가미 시즈카",
					},
				}

			}

			t = loadHTML("404")
		})

		w.WriteHeader(http.StatusNotFound)

		var x int
		if len(sayings) > 0 {
			ra := rand.New(rand.NewSource(time.Now().UnixNano()))
			x = ra.Intn(len(sayings))
		} else {
			x = 0
		}

		data := struct {
			Path string
			Saying
		}{
			r.URL.Path,
			sayings[x],
		}

		err := t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
	}
}
