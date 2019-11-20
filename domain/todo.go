package domain

import "time"

type Todo struct {
	Title       string
	Content     string
	createdDate time.Time
}
