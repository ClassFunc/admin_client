# Singleton Firebase Admin Client

## Install:

```sh
go get -u github.com/ClassFunc/admin_client
```

## Usage:

```go
package main

import (
    "github.com/ClassFunc/admin_client"
)

func main() {

    admin_client.New(
    "PROJECT_ID",
    "SERVICE_ACCOUNT_FILE_PATH",
    )
    // or more simpler
    admin_client.WithCredentialsFile("SERVICE_ACCOUNT_FILE_PATH")
    
    // then you can use
    admin_client.DB.Collection(...)
    admin_client.Auth.CreateUser(...)
    admin_client.DefaultBucket.Object("media/" + fileName)
    admin_client.Storage.Bucket("name").Object("media/" + fileName)
	
	
    // settings firestore client
    admin_client.EnvCollPrefix = "GO_COL_"
    fmt.Println(admin_client.CollPathFromEnv("GO_COL_USERS"))
    
    // actions with firestore collections
    admin_client.CollRef("GO_COL_USERS" /* or "users"*/)
    admin_client.CollGetAllDocs("users")
    admin_client.CollGetAllDocsData("users")
    
    // actions with firestore documents
    admin_client.DocRef("users", "123")
    admin_client.GetDoc("users", "123")
    admin_client.GetDocData("users", "123")
    
    type User struct {
        Name string
    }
    var adam = new(User)
    admin_client.GetDocDataTo("users", "123", adam)
}
```

## Lisence:
MIT

## Authors:
ClassFunc Softwares JSC

Website: https://classfunc.com
