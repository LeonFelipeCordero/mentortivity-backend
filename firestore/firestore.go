package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirestoreClient() *firestore.Client {
	app := newFirebaseApp()

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection established to firestore database")

	return client
}

func newFirebaseApp() *firebase.App {
	serviceAccount := option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT_FILE"))

	app, err := firebase.NewApp(context.Background(), nil, serviceAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection established to firebase project mentortivity")

	return app
}
