package common

import "cloud.google.com/go/firestore"

// DocDataWithId returns a doc with appended `_id` field if it's not already there.
func DocDataWithId(doc *firestore.DocumentSnapshot) map[string]interface{} {
	data := doc.Data()
	if _, ok := data["_id"]; !ok {
		data["_id"] = doc.Ref.ID
	}
	return data
}
