package report

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"github.com/joho/godotenv"
)

func TestReportGeneration(t *testing.T) {
	godotenv.Load("../.env.test.local")

	t.Run("Create Report", func(t *testing.T) {
		client := firestore.NewFirestoreClient()
		createNotebook()

		reportGenerator := DefaultReportGenerator{}
		got := reportGenerator.CreateReport(*client, schema.FullNotebook{
			Id: "12345",
			Notebook: schema.Notebook{
				Email:         "test@email.com",
				Interruptions: 1,
				Pomodoros:     3,
				PomodoroTimer: 4,
			},
		})
		want := schema.Report{Interruptions: 1, Pomodoros: 3, PomodoroTimer: 4, Date: time.Now().Format("01-02-2006")}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
		client.Close()
	})

	t.Run("Save report as sub collecion", func(t *testing.T) {
		client := firestore.NewFirestoreClient()
		createNotebook()

		reportGenerator := DefaultReportGenerator{}
		got := reportGenerator.CreateReport(*client, schema.FullNotebook{
			Id: "12345",
			Notebook: schema.Notebook{
				Email:         "test@email.com",
				Interruptions: 1,
				Pomodoros:     3,
				PomodoroTimer: 4,
			},
		})
		document, err := client.Collection("notebooks").Doc("12345").Collection("reports").Where("pomodoros", "==", 3).Documents(context.Background()).Next()

		if err != nil {
			t.Errorf("Could not retrive report docuement")
		}

		want := schema.Report{}
		document.DataTo(&want)

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
	client := firestore.NewFirestoreClient()
	client.Collection("notebook").Doc("12345").Create(context.Background(), map[string]interface{}{
		"email":         "test@email.com",
		"interruptions": 1,
		"pomodoros":     3,
		"pomodoroTimer": 4,
	})
	client.Close()
}
