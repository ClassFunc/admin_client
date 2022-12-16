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

	admin_client.EnvCollPrefix = "GO_COL_"
	fmt.Println(admin_client.CollPathFromEnv("GO_COL_USERS"))

	admin_client.CollRef("GO_COL_USERS" /* or "users"*/)
	admin_client.CollGetAllDocs("users")
	admin_client.CollGetAllDocsData("users")

	admin_client.DocRef("users", "123")
	admin_client.GetDoc("users", "123")
	admin_client.GetDocData("users", "123")

	type User struct {
		Name string
	}
	var adam = new(User)
	admin_client.GetDocDataTo("users", "123", adam)

}
