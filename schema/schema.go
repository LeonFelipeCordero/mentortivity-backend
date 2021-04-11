package schema

type Notebook struct {
	Email         string `firestore:"email,omitempty"`
	Interruptions int    `firestore:"interruptions,omitempty"`
	Pomodoros     int    `firestore:"pomodoros,omitempty"`
	PomodoroTimer int    `firestore:"pomodoroTimer,omitempty"`
}

type Task struct {
	Datetime    string `firestore:"datetime,omitempty"`
	Description string `firestore:"description,omitempty"`
	Status      string `firestore:"status,omitempty"`
}

type FullNotebook struct {
	Id       string
	Notebook Notebook
	Tasks    []Task
}

type Report struct {
	Interruptions int    `firestore:"interruptions,omitempty"`
	Pomodoros     int    `firestore:"pomodoros,omitempty"`
	PomodoroTimer int    `firestore:"pomodoroTimer,omitempty"`
	Date          string `firestore:"Date,omitempty"`
}
