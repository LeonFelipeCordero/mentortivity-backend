package notebook

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	gfirestore "cloud.google.com/go/firestore"
	internal "github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"github.com/joho/godotenv"
)

func TestNotebooksCollections(t *testing.T) {
	godotenv.Load("../.env.test.local")

	t.Run("Get notebook by id", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		createNotebook()

		got := getNotebookById(*client, "12345")
		want := schema.Notebook{Email: "test@email.com", Interruptions: 1, Pomodoros: 3, PomodoroTimer: 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		deleteData()
		client.Close()
	})

	t.Run("Clear notebook when closing day", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		createNotebook()

		reportGenerator := TestReportGenerator{reports: 0}
		CloseDay(*client, &reportGenerator, "12345")

		got := getNotebookById(*client, "12345")
		want := schema.Notebook{Email: "test@email.com", Interruptions: 0, Pomodoros: 0, PomodoroTimer: 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		deleteData()
		client.Close()
	})

	t.Run("Return report when closing day", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		createNotebook()

		reportGenerator := TestReportGenerator{reports: 0}
		got := CloseDay(*client, &reportGenerator, "12345")
		want := schema.Report{Interruptions: 1, Pomodoros: 3, PomodoroTimer: 4, Date: time.Now().Format("01-02-2006")}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		deleteData()
		client.Close()
	})
}

func deleteData() {
	client := http.Client{}
	request, err := http.NewRequest(
		"DELETE", "http://localhost:8080/emulator/v1/projects/mentortivity-96f2a/databases/(default)/documents", nil,
	)

	if err != nil {
		panic("Something went wrong creating request to delete documents")
	}

	resp, err := client.Do(request)

	if err != nil {
		panic("Something went wrong calling firestore to delete documentes")
	}

	log.Printf("Status: %s", resp.Status)
}

func createNotebook() {
	client := internal.NewFirestoreClient()
	client.Collection(notebookCollection).Doc("12345").Create(context.Background(), map[string]interface{}{
		"email":         "test@email.com",
		"interruptions": 1,
		"pomodoros":     3,
		"pomodoroTimer": 4,
	})
	client.Close()
}

type TestReportGenerator struct {
	reports int
}

func (reportGenerator *TestReportGenerator) CreateReport(client gfirestore.Client, notebook schema.FullNotebook) schema.Report {
	reportGenerator.reports += 1
	report := schema.Report{
		Interruptions: notebook.Notebook.Interruptions,
		Pomodoros:     notebook.Notebook.Pomodoros,
		PomodoroTimer: notebook.Notebook.PomodoroTimer,
		Date:          time.Now().Format("01-02-2006"),
	}

	return report
}
