package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string
	Description string
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

/*
{
	"complete": true
}
*/

type CompleteDTO struct {
	Complete bool
}

func (e ErrorDTO) ToJSON() string {
	d, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(d)
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("Empty title")
	}

	if t.Description == "" {
		return errors.New("Empty task description")
	}

	return nil
}
