package testingUtils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	internal "github.com/LeonFelipeCordero/mentortivity-backend/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"github.com/joho/godotenv"
)

func LoadTestEnv() {
	godotenv.Load("../.env.test.local")
}

func DeleteData() {
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

func CreateNotebook() {
	notebook := schema.Notebook{}
	LoadFromFiles("../testingUtils/notebook.json", &notebook)
	client := internal.NewFirestoreClient()
	client.Collection("notebooks").Doc("12345").Create(context.Background(), notebook)
	client.Close()
}

func CreateTasks() {
	tasks := []schema.Task{}
	LoadFromFiles("../testingUtils/task.json", &tasks)

	client := internal.NewFirestoreClient()

	collectionRef := client.Collection("notebooks").Doc("12345").Collection("tasks")

	for _, task := range tasks {
		log.Print(task)
		collectionRef.NewDoc().Set(context.Background(), task)
	}

	client.Close()
}

func LoadFromFiles(file string, target interface{}) {
	path, _ := filepath.Abs(file)
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Something wrong happened trying to read file => %v", err)
	}

	err = json.Unmarshal([]byte(data), &target)

	if err != nil {
		log.Fatalf("Something wrong happened serializing file %v", err)
	}
}

func LoadMockTasks() []schema.Task {
	tasks := []schema.Task{}
	LoadFromFiles("../testingUtils.json", tasks)
	return tasks
}
