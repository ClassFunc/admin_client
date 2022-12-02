# Sigleton Firebase Admin Client

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
  
  // then you can use
  admin_client.DB.Collection(...)
  admin_client.Auth.CreateUser(...)
  admin_client.DefaultBucket.Object("media/" + fileName)
  admin_client.Storage.Bucket("name").Object("media/" + fileName)
}
```

## Lisence:
MIT

## Authors:
ClassFunc Softwares JSC

Website: https://classfunc.com
