package models

import (
	"github.com/go-pg/pg"
)

type voteResult struct {
	IsVoteUp bool
	Value    int
}

func (vote *Vote) Create(db *pg.DB) (err error) {
	err = db.Insert(vote)
	return
}

func GetVotesByPostID(db *pg.DB, id int) (upVote int, downVote int, err error) {
	var result []voteResult
	err = db.Model(&Vote{}).Column("is_vote_up").
		ColumnExpr("count(*) as value").
		Where("post_id = ?", id).
		Group("is_vote_up").
		Select(&result)

	if len(result) == 2 {

		if result[0].IsVoteUp {
			upVote = result[0].Value
			downVote = result[1].Value
		} else {
			downVote = result[0].Value
			upVote = result[1].Value
		}

	} else if len(result) == 1 {

		if result[0].IsVoteUp {
			upVote = result[0].Value
		} else {
			downVote = result[0].Value
		}

	}

	return
}

func VotePost(db *pg.DB, postID int, sess *Session, isVoteUp bool) (voted int, err error) {
	// Check if user voted
	wasVoteUp, err := checkIfVoted(db, postID, sess)

	if isVoteUp {
		voted = 1 // UpVote
	} else {
		voted = 2 // DownVote
	}

	if err != nil { // If user didn't vote

		vote := Vote{
			SessionID: sess.ID,
			Session:   sess,
			PostID:    postID,
			IsVoteUp:  isVoteUp,
		}

		err = vote.Create(db)
		return

	} else if wasVoteUp == isVoteUp {

		_, err = db.Model(&Vote{}).
			Where("post_id = ?", postID).
			Where("session_id = ?", sess.ID).
			Delete()

		voted = 0 // None
		return

	} else {

		_, err = db.Model(&Vote{}).Set("is_vote_up = ?", isVoteUp).
			Where("post_id = ?", postID).
			Where("session_id = ?", sess.ID).
			Update()

		return

	}

}

func checkIfVoted(db *pg.DB, postID int, sess *Session) (isVoteUp bool, err error) {
	err = db.Model(&Vote{}).
		Column("is_vote_up").
		Where("session_id = ?", sess.ID).
		Where("post_id = ?", postID).
		Select(&isVoteUp)

	return
}

func GetWhatUserVotedFromPosts(db *pg.DB, ps *[]Post, sess *Session) {
	for i := range *ps {
		isVoteUp, err := checkIfVoted(db, (*ps)[i].ID, sess)
		if err != nil {
			(*ps)[i].WhatUserVoted = 0
		} else if isVoteUp {
			(*ps)[i].WhatUserVoted = 1
		} else {
			(*ps)[i].WhatUserVoted = 2
		}
	}
}
