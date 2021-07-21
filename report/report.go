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

	doneTasks := filter	

	report := schema.Report{
		Interruptions:      notebook.Notebook.Interruptions,
		Pomodoros:          notebook.Notebook.Pomodoros,
		PomodoroTimer:      notebook.Notebook.PomodoroTimer,
		Date:               time.Now().Format("01-02-2006"),
		PendingToDoneRatio: 0,
		WeakVerbs:          []string{},
	}

	client.Collection(notebookCollection).Doc(notebook.Id).Collection(reportCollection).NewDoc().Create(context.Background(), report)

	return report
}
