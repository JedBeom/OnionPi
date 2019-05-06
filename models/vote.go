package models

import (
	"github.com/go-pg/pg"
)

type voteResult struct {
	IsVoteUp bool
	Value    int
}

func GetVotesByPostID(db *pg.DB, id int) (upVote int, downVote int, err error) {
	var result []voteResult
	err = db.Model(&Vote{}).Column("is_vote_up").ColumnExpr("count(*) as value").
		Where("post_id = ?", id).Group("is_vote_up").Select(&result)

	if len(result) == 2 {
		upVote = result[0].Value
		downVote = result[1].Value
	} else if len(result) == 1 {

		if result[0].IsVoteUp {
			upVote = result[0].Value
		} else {
			downVote = result[0].Value
		}

	}

	return
}
