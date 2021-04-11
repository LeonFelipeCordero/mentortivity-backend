package report

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
)

const notebookCollection = "notebooks"
const reportCollection = "reports"

type ReportGenerator interface {
	CreateReport(client firestore.Client, notebook schema.FullNotebook) schema.Report
}

type DefaultReportGenerator struct {
}

func (reportGenerator *DefaultReportGenerator) CreateReport(client firestore.Client, notebook schema.FullNotebook) schema.Report {

	client.Collection(notebookCollection).Doc(notebook.Id).Collection(reportCollection).NewDoc().Create(context.Background(), map[string]interface{}{
		"date":          time.Now().Format("01-02-2006"),
		"pomodoros":     notebook.Notebook.Pomodoros,
		"interruptions": notebook.Notebook.Interruptions,
		"pomodoroTimer": notebook.Notebook.PomodoroTimer,
	})

	return schema.Report{
		Interruptions: notebook.Notebook.Interruptions,
		Pomodoros:     notebook.Notebook.Pomodoros,
		PomodoroTimer: notebook.Notebook.PomodoroTimer,
		Date:          time.Now().Format("01-02-2006"),
	}
}
