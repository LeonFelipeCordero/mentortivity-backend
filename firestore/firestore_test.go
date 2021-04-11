package firestore

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestFiresotreClient(t *testing.T) {

	t.Run("Create firestore client", func(t *testing.T) {
		godotenv.Load("../.env.test.local")
		got := NewFirestoreClient()

		if got == nil {
			t.Errorf("Error while creating firestore client")
		}

		got.Close()
	})
}
