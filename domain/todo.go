package domain

type Todo struct {
	TodoId  int64  `json:"todo_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
