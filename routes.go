package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/go-chi/chi/middleware"

	"github.com/JedBeom/onionpi/models"
	"github.com/go-chi/chi"
)

func route() (c *chi.Mux) {
	c = chi.NewMux()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Use(checkCookie)

	notFound := notFoundHandler()
	c.NotFound(notFound)
	c.Get("/static/*", staticHandler(notFound))

	c.Get("/", index())
	c.Get("/vote/{postID}/{voting}", vote())
	c.Post("/new/post", submitPost())

	return
}

func submitPost() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		content := r.PostFormValue("content")

		if len(strings.Fields(content)) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sess := getOrCreateCookie(db, w, r)
		post := models.Post{
			SessionID: sess.ID,
			Session:   sess,
			Content:   content,
		}

		err := post.Create(db)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)

	}
}

func index() http.HandlerFunc {
	var (
		init sync.Once
		t    *template.Template
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			t = loadHTML("base", "main", "posts")
		})
		posts, err := models.GetPosts(db)
		if err != nil {
			_, _ = w.Write([]byte("Error"))
			return
		}

		sess := getOrCreateCookie(db, w, r)

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

		fmt.Println(models.GetVotesByPostID(db, 1))

	}

}

func vote() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "postID")
		id, _ := strconv.Atoi(idStr)
		p, err := models.GetPostByID(db, id)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		voting := chi.URLParam(r, "voting")
		if voting == "+" {
			err = p.VoteUp(db)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else if voting == "-" {
			err = p.VoteDown(db)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}