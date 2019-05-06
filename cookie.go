package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-pg/pg"

	"github.com/JedBeom/onionpi/models"
)

const (
	SessionCookie = "_yayoiori"
)

func checkCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie(SessionCookie)
		if err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}

		if c.Value == "" {
			log.Println("No value")
			next.ServeHTTP(w, r)
			return
		}

		sess, err := models.SessionByUUID(db, c.Value)
		if err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "sess", sess)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getOrCreateCookie(db *pg.DB, w http.ResponseWriter, r *http.Request) *models.Session {
	sess, ok := r.Context().Value("sess").(*models.Session)
	if !ok || sess == nil {
		var err error
		sess, err = models.NewSession(db, r)
		if err != nil {
			log.Println(err)
			return sess
		}

		cookie := &http.Cookie{
			Name:  SessionCookie,
			Value: sess.UUID,
		}
		http.SetCookie(w, cookie)
	}

	return sess
}
