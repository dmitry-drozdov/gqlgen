// Code generated by github.com/dmitry-drozdov/gqlgen, DO NOT EDIT.

package deferexample

type Mutation struct {
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Query struct {
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	User   *User  `json:"user"`
	userID string `json:"-"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
