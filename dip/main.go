package main

import "log"

type SQLHandler interface {
	Execute()
}

type sqlHandler struct{}

func (s *sqlHandler) Execute() {
	log.Println("SQL Execute")
}

func NewSQLHandler() SQLHandler {
	log.Println("New SQL Handler")
	return &sqlHandler{}
}

type FooRepository struct {
	SQLHandler SQLHandler
}

func (f *FooRepository) Find() {
	f.SQLHandler.Execute()
}

// Before
// FooRepository -> sqlHandler
// After
// FooRepository -> SQLHandler Interface <- sqlHandler

func main() {
	sql := NewSQLHandler()
	foo := &FooRepository{sql}
	foo.Find()
}
