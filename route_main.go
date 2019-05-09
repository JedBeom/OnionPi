package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/JedBeom/onionpi/models"
)

func ShowMain() http.HandlerFunc {
	var (
		init sync.Once
		t    *template.Template
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			t = loadHTML("base", "main", "li")
		})
		sess := getOrCreateCookie(db, w, r)

		posts, err := models.GetPosts(db, sess, 10)
		if err != nil {
			_, _ = w.Write([]byte("Error"))
			return
		}

		data := struct {
			Posts      []models.Post
			HasSession bool
		}{
			posts,
			sess != nil,
		}
		err = t.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Println(err)
		}

	}

}
