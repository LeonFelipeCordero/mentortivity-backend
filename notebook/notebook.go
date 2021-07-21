package notebook

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/LeonFelipeCordero/mentortivity-backend/report"
	"github.com/LeonFelipeCordero/mentortivity-backend/schema"
)

const notebookCollection = "notebooks"

func CloseDay(client firestore.Client, reportGenerator report.ReportGenerator, id string) schema.Report {
	notebook := getNotebookById(client, id)

	report := reportGenerator.CreateReport(client,
		schema.FullNotebook{
			Id:       id,
			Notebook: notebook,
		},
	)

	resetCounters(client, id)
	// TODO send tasks to expired after closing day

	return report
}

func getNotebookById(client firestore.Client, id string) schema.Notebook {
	document, err := client.Collection(notebookCollection).Doc(id).Get(context.Background())

	if err != nil {
		log.Printf("error: %v", err)
		panic(fmt.Sprintf("Impossible to get document in collection `%s` with id `%s`", notebookCollection, id))
	}

	notebook := schema.Notebook{}
	document.DataTo(&notebook)
	return notebook
}

func resetCounters(client firestore.Client, id string) {
	client.Collection(notebookCollection).Doc(id).Update(context.Background(), []firestore.Update{
		{Path: "pomodoros", Value: 0},
		{Path: "interruptions", Value: 0},
		{Path: "pomodoroTimer", Value: 0},
	})
}
