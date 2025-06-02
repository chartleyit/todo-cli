package models

import "time"

type Status int

const (
	ToDo       Status = 0
	Inprogress Status = 1
	Done       Status = 2
)

type TodoItem struct {
	Id          int       `json:"id"`
	ParentId    int       `json:"parent"`
	ChildrenIds []int     `json:"childrend"`
	Task        string    `json:"task"`
	CreatedAt   time.Time `json:"created"`
	Due         time.Time `json:"due"`
	Status      Status    `json:"status"`
}

func (s Status) String() string {
	return [...]string{"ToDo", "InProgress", "Done"}[s]
}
