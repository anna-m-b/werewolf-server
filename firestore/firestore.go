package firestore

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/anna-m-b/werewolf-server/entity"
	"google.golang.org/api/option"
)

func init() {
  sa := option.WithCredentialsFile("./ServiceAccount.json")
  app, err := firebase.NewApp(context.Background(), nil, sa)
  if err != nil {
    log.Fatalf("Failed to set up Firebase app: %v", err)
  }

  db, err := app.Firestore(context.Background())
  if err != nil {
    log.Fatalf("Failed to access Firestore: %v", err)
  }

  _, _, err = db.Collection("games").Add(context.Background(), entity.CreateGame(entity.CreatePlayer("Normie")))
  if err != nil {
     log.Fatalf("Failed to create a Game document: %v", err)
  }
  
  defer db.Close()
}
