package models

type Student struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}
