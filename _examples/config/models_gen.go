// Code generated by github.com/dmitry-drozdov/gqlgen, DO NOT EDIT.

package config

type Mutation struct {
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Query struct {
}

type Todo struct {
	ID          string    `json:"id"`
	DatabaseID  int       `json:"databaseId"`
	Description string    `json:"text"`
	Done        bool      `json:"done"`
	User        *User     `json:"user"`
	Query       *Query    `json:"query"`
	Mutation    *Mutation `json:"mutation"`
}
