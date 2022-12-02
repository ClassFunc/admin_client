package admin_client

import (
	"context"
	"firebase.google.com/go/v4/auth"
)

var Auth *auth.Client

// GetAuth initializes the firebase auth client
func GetAuth() (client *auth.Client, err error) {
	client, err = App.Auth(context.Background())
	return
}
