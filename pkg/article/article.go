// article.go

// Package article defines the structure of the article
// and the related reporter that creates the article
// and publisher that outputs the article
package article

import (
	// "strings"
	"time"
)

// struct Article defines properties
type Article struct {
	title, description string
	timeStamp          time.Time
}

// function Title returns Title string
func (art *Article) Title() string {
	return art.title
}

// function Description returns Description string
func (art *Article) Description() string {
	return art.description
}

// function TimeStamp returns date time created
func (art *Article) TimeStamp() time.Time {
	return art.timeStamp
}

// func New creates new Article
func New(Title, Description string) Article {
	var this Article

	// init properties
	this.title = Title
	this.description = Description
	this.timeStamp = time.Now()

	return this
}

// interface Reporter creates Articles
type Reporter interface {
	Report(Title string) Article
}

// interface Publisher outputs Articles
type Publisher interface {
	Publish(myArticle Article) error
	MustPublish(myArticle Article)
}
