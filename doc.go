package admin_client

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/ClassFunc/admin_client/common"
)

// DocRef returns a reference to the document with the given name.
func DocRef(coll, docId string) *firestore.DocumentRef {
	realCollName := CollPathFromEnv(coll)
	return DB.Collection(realCollName).Doc(docId)
}

// GetDoc   returns the document with the given name.
func GetDoc(coll, docId string) (*firestore.DocumentSnapshot, error) {
	ctx := context.Background()
	doc, err := DocRef(coll, docId).Get(ctx)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// GetDocData returns the document data with the given name.
func GetDocData(coll, docId string) (map[string]interface{}, error) {
	ctx := context.Background()
	doc, err := DocRef(coll, docId).Get(ctx)
	if err != nil {
		return nil, err
	}
	return common.DocDataWithId(doc), nil
}

func GetDocDataTo(coll, docId string, ptr interface{}) error {
	ctx := context.Background()
	doc, err := DocRef(coll, docId).Get(ctx)
	if err != nil {
		return err
	}
	return doc.DataTo(ptr)
}
