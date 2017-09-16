package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Comments_20170916_222233 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Comments_20170916_222233{}
	m.Created = "20170916_222233"

	migration.Register("Comments_20170916_222233", m)
}

// Run the migrations
func (m *Comments_20170916_222233) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comments(id SERIAL PRIMARY KEY, comments  TEXT NOT NULL, id_user SERIAL REFERENCES users(id), id_article SERIAL REFERENCES articles(id), created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, updated_at timestamp NOT NULL)")
}

// Reverse the migrations
func (m *Comments_20170916_222233) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE comments")
}