package report

import (
	"context"
	"reflect"
	"testing"
	"time"

	gfirestore "cloud.google.com/go/firestore"
	internal "github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"github.com/LeonFelipeCordero/mentortivity-backend/testingUtils"
)

func TestReportGeneration(t *testing.T) {
	testingUtils.LoadTestEnv()

	t.Run("Create Report", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		got := createReport(client)
		want := schema.Report{
			Interruptions:      1,
			Pomodoros:          3,
			PomodoroTimer:      4,
			Date:               time.Now().Format("01-02-2006"),
			PendingToDoneRatio: "3:6",
			WeakVerbs:          []string{"prepare", "create", "implement", "investigate"},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		testingUtils.DeleteData()
	})

	t.Run("Save report as sub collecion", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		got := createReport(client)
		document, err := client.Collection("notebooks").Doc("12345").Collection("reports").Where("pomodoros", "==", 3).Documents(context.Background()).Next()

		if err != nil {
			t.Errorf("Could not retrive report docuement")
		}

		want := schema.Report{}
		document.DataTo(&want)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		testingUtils.DeleteData()
	})

	t.Run("Should get done pending ratio from tasks", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		got := createReport(client).PendingToDoneRatio
		want := "3:6"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		testingUtils.DeleteData()
	})

	t.Run("Should get done pending ratio from tasks", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()

		got := createReport(client).WeakVerbs
		want := []string{"prepare", "create", "implement", "investigate"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		testingUtils.DeleteData()
	})

	testingUtils.DeleteData()
}

func createReport(client *gfirestore.Client) schema.Report {
	reportGenerator := DefaultReportGenerator{}
	return reportGenerator.CreateReport(*client, schema.FullNotebook{
		Id: "12345",
		Notebook: schema.Notebook{
			Email:         "test@email.com",
			Interruptions: 1,
			Pomodoros:     3,
			PomodoroTimer: 4,
		},
		Tasks: testingUtils.LoadMockTasks(),
	})

}
