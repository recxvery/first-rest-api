package todo


import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time //поле является указателем, т.к. в json если мы не выполнили задачу будет равен null, что для клиента будет удобно.
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completeTime
}


func (t *Task) Uncomplete() {
	completeTime := time.Now()

	t.Completed = false
	t.CompletedAt = &completeTime
}
