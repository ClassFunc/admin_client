package main

import (
	"fmt"
	"github.com/ClassFunc/admin_client"
)

const serviceAccountFilePath = "example/YOUR_SERVICE_ACCOUNT_KEY.json"

func main() {
	admin_client.WithCredentialsFile(serviceAccountFilePath)
	fmt.Println(
		admin_client.App,
		admin_client.Auth,
		admin_client.Firestore,
		admin_client.Storage,
		admin_client.DefaultBucket,
	)
}
