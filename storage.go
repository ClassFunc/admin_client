package admin_client

import (
	cloudStorage "cloud.google.com/go/storage"
	"context"
	"firebase.google.com/go/v4/storage"
)

var Storage *storage.Client
var DefaultBucket *cloudStorage.BucketHandle

// GetStorageClient returns the firebase storage client
func GetStorageClient() (client *storage.Client, err error) {
	client, err = App.Storage(context.Background())
	return
}
