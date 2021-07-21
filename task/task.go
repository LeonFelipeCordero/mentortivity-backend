package task

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
	"google.golang.org/api/iterator"
)

type TaskLoader interface {
	LoadTasks(client *firestore.Client, id string) []schema.Task
}

type DefaultTaskLoader struct{}

func (loader *DefaultTaskLoader) LoadTasks(client *firestore.Client, id string) []schema.Task {
	tasks := []schema.Task{}

	tasksRefs := client.Collection("notebooks").Doc(id).Collection("tasks").Documents(context.Background())
	for {
		doc, err := tasksRefs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(fmt.Sprintf("Something went wrong trying to get tasks for document %s, %v", id, err))
		}
		task := schema.Task{}
		doc.DataTo(&task)
		tasks = append(tasks, task)
	}

	return tasks
}
