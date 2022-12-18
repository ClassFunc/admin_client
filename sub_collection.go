package admin_client

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/ClassFunc/admin_client/common"
	"github.com/samber/lo"
)

// SubCollRef returns a reference to the sub-collection with the given name.
func SubCollRef(coll, docId, subColl string) *firestore.CollectionRef {
	return DocRef(coll, docId).Collection(subColl)
}

// SubCollGetAllDocs returns all documents in the sub-collection.
func SubCollGetAllDocs(coll, docId, subColl string) ([]*firestore.DocumentSnapshot, error) {
	ctx := context.Background()
	snap, err := SubCollRef(coll, docId, subColl).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	return snap, nil
}

// SubCollGetAllDocsData returns all documents data in the sub-collection.
func SubCollGetAllDocsData(coll, docId, subColl string) ([]map[string]interface{}, error) {
	snap, err := SubCollGetAllDocs(coll, docId, subColl)
	if err != nil {
		return nil, err
	}
	res := lo.Map(snap, func(doc *firestore.DocumentSnapshot, index int) map[string]interface{} {
		return common.DocDataWithId(doc)
	})
	return res, nil
}

// SubDocRef returns a reference to the document with the given name in the sub-collection.
func SubDocRef(coll, docId, subColl, subDocId string) *firestore.DocumentRef {
	return SubCollRef(coll, docId, subColl).Doc(subDocId)
}

// GetSubDoc returns the document with the given name in the sub-collection.
func GetSubDoc(coll, docId, subColl, subDocId string) (*firestore.DocumentSnapshot, error) {
	ctx := context.Background()
	return SubDocRef(coll, docId, subColl, subDocId).Get(ctx)
}

// GetSubDocData returns the document data with the given name in the sub-collection.
func GetSubDocData(coll, docId, subColl, subDocId string) (map[string]interface{}, error) {
	doc, err := GetSubDoc(coll, docId, subColl, subDocId)
	if err != nil {
		return nil, err
	}
	return common.DocDataWithId(doc), nil
}
