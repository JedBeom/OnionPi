package models

import "github.com/go-pg/pg"

func (p *Post) Create(db *pg.DB) (err error) {
	err = db.Insert(p)
	return
}

/*
func GetPostByID(db *pg.DB, id int) (p Post, err error) {
	p.ID = id
	err = db.Select(&p)
	return
}
*/

func GetPosts(db *pg.DB, sess *Session, limit int) (posts []Post, err error) {
	err = db.Model(&posts).Order("id DESC").Limit(limit).Select()

	// Calculate TotalVote
	for i := range posts {
		up, down, err := GetVotesByPostID(db, posts[i].ID)
		if err != nil {
			continue
		}
		posts[i].TotalVote = up - down
	}

	if sess != nil {
		GetWhatUserVotedFromPosts(db, &posts, sess)
	}

	return
}

func GetPostByContent(db *pg.DB, content string) (p Post, err error) {
	err = db.Model(&p).Where("content = ?", content).First()
	return
}
