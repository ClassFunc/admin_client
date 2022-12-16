package admin_client

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/ClassFunc/admin_client/common"
	"github.com/samber/lo"
	"os"
	"strings"
)

var (
	EnvCollPrefix = "COLL_"
)

func CollPathFromEnv(envVarNameOrCollPath string) string {
	if strings.HasPrefix(envVarNameOrCollPath, EnvCollPrefix) {
		return os.Getenv(envVarNameOrCollPath)
	}
	collEnvVar := EnvCollPrefix + strings.ToUpper(envVarNameOrCollPath)
	collPath := os.Getenv(collEnvVar)
	if lo.IsEmpty(collPath) {
		collPath = envVarNameOrCollPath
	}
	return collPath
}

// CollRef returns a reference to the collection with the given name.
func CollRef(path string) *firestore.CollectionRef {
	realCollPath := CollPathFromEnv(path)
	return DB.Collection(realCollPath)
}

// CollGetAllDocs returns all documents in the collection.
func CollGetAllDocs(coll string) ([]*firestore.DocumentSnapshot, error) {
	ctx := context.Background()
	snap, err := CollRef(coll).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	return snap, nil
}

// CollGetAllDocsData returns all documents data in the collection.
func CollGetAllDocsData(coll string) ([]map[string]interface{}, error) {
	snap, err := CollGetAllDocs(coll)
	if err != nil {
		return nil, err
	}
	res := lo.Map(snap, func(doc *firestore.DocumentSnapshot, index int) map[string]interface{} {
		return common.DocDataWithId(doc)
	})
	return res, nil
}
