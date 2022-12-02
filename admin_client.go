package admin_client

import (
	"sync"
)

// firebase instance
var (
	once sync.Once
)

func New(projectId, serviceAccountFilePath string) {
	var err error

	singletonApp(projectId, serviceAccountFilePath)

	// App
	if App == nil {
		panic("App is nil")
	}

	// Auth
	Auth, err = GetAuth()
	if err != nil {
		panic(err)
	}

	// Firestore, DB
	Firestore, err = GetFirestoreClient()
	if err != nil {
		panic(err)
	}
	DB = Firestore

	// Storage, DefaultBucket
	Storage, err = GetStorageClient()
	if err != nil {
		panic(err)
	}
	DefaultBucket, err = Storage.DefaultBucket()
	if err != nil {
		panic(err)
	}
}

// singletonApp returns the singleton instance
func singletonApp(projectId, saFilePath string) {
	once.Do(func() { // <-- atomic, does not allow repeating
		App = newApp(projectId, saFilePath)
	})
}
