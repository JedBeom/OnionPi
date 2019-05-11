package models

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"

	uuid "github.com/satori/go.uuid"
)

var (
	IsCloudflare = false
)

func NewSession(db *pg.DB, r *http.Request) (sess *Session, err error) {
	sess = &Session{}
	if !IsCloudflare {
		sess.IP = r.RemoteAddr
	} else {
		sess.IP = r.Header.Get("X-Forwarded-For")
	}

	a := strings.Split(sess.IP, ":")
	if len(a) == 2 {
		sess.IP = a[0]
	}

	sess.UserAgent = r.UserAgent()
	sess.UUID = uuid.NewV4().String()

	_, err = db.Model(sess).Returning("*").Insert()

	return
}

func SessionByUUID(db *pg.DB, value string) (sess *Session, err error) {
	sess = &Session{UUID: value}
	err = db.Model(sess).Where("uuid = ?", sess.UUID).Select()
	return
}
