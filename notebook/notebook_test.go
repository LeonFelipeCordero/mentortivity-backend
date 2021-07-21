package notebook

import (
	"reflect"
	"testing"
	"time"

	gfirestore "cloud.google.com/go/firestore"
	internal "github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"github.com/LeonFelipeCordero/mentortivity-backend/testingUtils"
)

func TestNotebooksCollections(t *testing.T) {
	testingUtils.LoadTestEnv()

	t.Run("Get notebook by id", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		got := getNotebookById(*client, "12345")
		want := schema.Notebook{Email: "test@email.com", Interruptions: 1, Pomodoros: 3, PomodoroTimer: 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		testingUtils.DeleteData()
		client.Close()
	})

	t.Run("Clear notebook when closing day", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		reportGenerator := TestReportGenerator{reports: 0}
		CloseDay(*client, &reportGenerator, "12345")

		got := getNotebookById(*client, "12345")
		want := schema.Notebook{Email: "test@email.com", Interruptions: 0, Pomodoros: 0, PomodoroTimer: 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		testingUtils.DeleteData()
		client.Close()
	})

	t.Run("Return report when closing day", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		reportGenerator := TestReportGenerator{reports: 0}
		got := CloseDay(*client, &reportGenerator, "12345")
		want := schema.Report{Interruptions: 1, Pomodoros: 3, PomodoroTimer: 4, Date: time.Now().Format("01-02-2006")}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		client.Close()
	})

	t.Run("Load tasks per report", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		reportGenerator := TestReportGenerator{reports: 0}
		got := CloseDay(*client, &reportGenerator, "12345")
		want := schema.Report{Interruptions: 1, Pomodoros: 3, PomodoroTimer: 4, Date: time.Now().Format("01-02-2006")}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		client.Close()
	})

	testingUtils.DeleteData()
}

type TestReportGenerator struct {
	reports int
}

func (reportGenerator *TestReportGenerator) CreateReport(client gfirestore.Client, notebook schema.FullNotebook) schema.Report {
	reportGenerator.reports += 1
	report := schema.Report{
		Interruptions:      notebook.Notebook.Interruptions,
		Pomodoros:          notebook.Notebook.Pomodoros,
		PomodoroTimer:      notebook.Notebook.PomodoroTimer,
		Date:               time.Now().Format("01-02-2006"),
		PendingToDoneRatio: 0,
		WeakVerbs:          []string{},
	}
	return report
}

type TestTaskLoader struct {
	loads int
}

func (loader *TestTaskLoader) LoadTasks(client gfirestore.Client, id string) []schema.Task {
	loader.loads += 1
	tasks := []schema.Task{}
	testingUtils.LoadFromFiles("../testUtils/tasks.json", tasks)
	return tasks
}
