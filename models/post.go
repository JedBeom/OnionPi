package models

import "github.com/go-pg/pg"

func (p *Post) Create(db *pg.DB) (err error) {
	err = db.Insert(p)
	return
}

func (p *Post) VoteUp(db *pg.DB) (err error) {
	//p.UpVote += 1
	err = db.Update(p)
	return
}

func (p *Post) VoteDown(db *pg.DB) (err error) {
	//p.DownVote += 1
	err = db.Update(p)
	return
}

func GetPostByID(db *pg.DB, id int) (p Post, err error) {
	p.ID = id
	err = db.Select(&p)
	return
}

func GetPosts(db *pg.DB) (posts []Post, err error) {
	err = db.Model(&posts).Order("id ASC").Select()

	// Calculate TotalVote
	for i := range posts {
		a, b, err := GetVotesByPostID(db, posts[i].ID)
		if err != nil {
			continue
		}
		posts[i].TotalVote = a - b
	}

	return
}