package todo

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlradyExists = errors.New("Task already  exist")
 