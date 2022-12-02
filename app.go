package admin_client

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

var App *firebase.App

// newApp initializes the firebase client
func newApp(projectId, saFilePath string) *firebase.App {
	app, err := firebase.NewApp(
		context.Background(),
		appConf(projectId), appOptions(saFilePath)...,
	)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
		return nil
	}
	return app
}

func appConf(projectId string) *firebase.Config {
	conf := &firebase.Config{
		ProjectID:     projectId,
		StorageBucket: fmt.Sprintf("%s.appspot.com", projectId),
	}
	return conf
}

func appOptions(saFilePath string) []option.ClientOption {
	opt := option.WithCredentialsFile(saFilePath)
	return []option.ClientOption{opt}
}
