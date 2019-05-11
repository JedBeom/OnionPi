package main

import (
	"net/http"
	"strings"

	"github.com/JedBeom/onionpi/models"
)

func SubmitPost() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("잘못된 요청입니다."))
			return
		}

		content := r.Form.Get("content")
		if len(strings.Fields(content)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("내용이 없습니다."))
			return
		}

		if _, err := models.GetPostByContent(db, content); err == nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("중복된 내용입니다."))
			return
		}

		sess, err := models.SessionByUUID(db, r.Form.Get("cookie"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("유효하지 않은 세션입니다."))
			return
		}

		post := models.Post{
			SessionID: sess.ID,
			Session:   sess,
			Content:   content,
		}

		err = post.Create(db)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
		return

	}
}
