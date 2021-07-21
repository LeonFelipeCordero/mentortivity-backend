package schema

type Notebook struct {
	Email         string `firestore:"email,omitempty" json:"email"`
	Interruptions int    `firestore:"interruptions,omitempty" json:"interruptions"`
	Pomodoros     int    `firestore:"pomodoros,omitempty" json:"pomodoros"`
	PomodoroTimer int    `firestore:"pomodoroTimer,omitempty" json:"pomodoroTimer"`
}

type Task struct {
	Datetime    string `firestore:"datetime,omitempty" json:"datetinme"`
	Description string `firestore:"description,omitempty" json:"description"`
	Status      string `firestore:"status,omitempty" json:"status"`
}

type FullNotebook struct {
	Id       string
	Notebook Notebook
	Tasks    []Task
}

type Report struct {
	Interruptions      int      `firestore:"interruptions,omitempty"`
	Pomodoros          int      `firestore:"pomodoros,omitempty"`
	PomodoroTimer      int      `firestore:"pomodoroTimer,omitempty"`
	PendingToDoneRatio string   `firestore:"pendingToDoneRatio,omitempty"`
	WeakVerbs          []string `firestore:"weakVerbs,omitempty"`
	Date               string   `firestore:"date,omitempty"`
}
