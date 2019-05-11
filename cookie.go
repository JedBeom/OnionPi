package main

import (
	"context"
	"net/http"
	"time"

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
			next.ServeHTTP(w, r)
			return
		}

		if c.Value == "" {
			next.ServeHTTP(w, r)
			return
		}

		sess, err := models.SessionByUUID(db, c.Value)
		if err != nil {
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
			return sess
		}

		cookie := &http.Cookie{
			Name:    SessionCookie,
			Value:   sess.UUID,
			Expires: time.Now().AddDate(0, 6, 0),
		}
		http.SetCookie(w, cookie)
	}

	return sess
}
