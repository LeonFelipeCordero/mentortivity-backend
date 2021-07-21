package task

import (
	"reflect"
	"sort"
	"testing"

	internal "github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"github.com/LeonFelipeCordero/mentortivity-backend/testingUtils"
)

func TestLoadTasks(t *testing.T) {
	testingUtils.LoadTestEnv()

	t.Run("Load tasks by notebook", func(t *testing.T) {
		client := internal.NewFirestoreClient()
		testingUtils.CreateNotebook()
		testingUtils.CreateTasks()

		taskLoader := DefaultTaskLoader{}
		got := taskLoader.LoadTasks(client, "12345")
		want := expectedTasks()

		sort.Slice(got, func(i, j int) bool {
			return got[i].Description < got[j].Description
		})
		sort.Slice(want, func(i, j int) bool {
			return want[i].Description < want[j].Description
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}

		testingUtils.DeleteData()
	})
}

func expectedTasks() []schema.Task {
	tasks := []schema.Task{}
	testingUtils.LoadFromFiles("../testingUtils/task.json", &tasks)
	return tasks
}
