package models

import (
	"time"
)

type Post struct {
	ID int

	SessionID int `sql:",pk"`
	Session   *Session
	Content   string `sql:",notnull"`

	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time

	WhatUserVoted int `sql:"-"` // go-pg ignores this field.
	TotalVote     int `sql:"-"` // go-pg ignores this field.
}

type Vote struct {
	ID int

	SessionID int `sql:",pk",sql:"unique:aa"`
	Session   *Session
	PostID    int `sql:",pk",sql:"unique:aa"`
	Post      *Post

	IsVoteUp bool `sql:",notnull"`

	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time
}

type Session struct {
	ID   int
	UUID string `sql:",unique"`

	IP        string
	UserAgent string

	CreatedAt time.Time `sql:"default:now()"`
	DeletedAt time.Time `pg:",soft_delete"`
}
