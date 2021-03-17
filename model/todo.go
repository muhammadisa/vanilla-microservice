package model

import uuid "github.com/satori/go.uuid"

// Singular
// Todo struct model
type Todo struct {
	ID uuid.UUID
	Name string
	Completed bool
}

// Plural
// Todos struct model
type Todos []Todo
