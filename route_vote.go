package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JedBeom/onionpi/models"
)

func VotePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := strconv.Atoi(r.Form.Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isVoteUp, err := SToB(r.Form.Get("is_vote_up"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sess := getOrCreateCookie(db, w, r)

		whatVoted := 0
		if isVoteUp {
			whatVoted, err = models.VotePost(db, id, sess, true)
		} else {
			whatVoted, err = models.VotePost(db, id, sess, false)
		}

		up, down, err := models.GetVotesByPostID(db, id)
		var updatedVote int
		if err == nil {
			updatedVote = up - down
		}

		data := struct {
			TotalVote int `json:"total_vote"`
			UserVote  int `json:"user_vote"`
		}{
			updatedVote,
			whatVoted,
		}

		j, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			_, _ = w.Write(j)
			w.WriteHeader(http.StatusOK)
		}

	}
}
