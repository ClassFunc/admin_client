package admin_client

import (
	"cloud.google.com/go/firestore"
	"context"
)

var Firestore *firestore.Client
var DB *firestore.Client

// GetFirestoreClient returns the firebase firestore client
func GetFirestoreClient() (client *firestore.Client, err error) {
	client, err = App.Firestore(context.Background())
	return
}
