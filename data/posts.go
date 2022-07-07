package data

import "time"

type Post struct {
	Slug     string
	Title    string
	Body     string
	PostedAt time.Time
}

var (
	Posts = []Post{
		{
			Slug:     "post-1",
			Title:    "Post 1",
			Body:     "This is the body of Post 1",
			PostedAt: time.Date(2022, 7, 5, 0, 0, 0, 0, time.Local),
		}, {
			Slug:     "post-2",
			Title:    "Post 2",
			Body:     "This is the body of Post 2",
			PostedAt: time.Date(2022, 7, 7, 0, 0, 0, 0, time.Local),
		},
	}
)
