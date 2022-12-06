package admin_client

import (
	"encoding/json"
	"os"
)

type ServiceAccount struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func WithCredentialsFile(filePath string) {
	sa := MustReadSAFile(filePath)
	New(sa.ProjectId, filePath)
}

func MustReadSAFile(path string) *ServiceAccount {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	var serviceAccount ServiceAccount
	err = decoder.Decode(&serviceAccount)
	if err != nil {
		panic(err)
	}
	return &serviceAccount
}
