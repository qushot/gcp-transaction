package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, firestore.DetectProjectID)
	if err != nil {
		return
	}
	defer client.Close()

	if err := client.RunTransaction(ctx, sampleFunc); err != nil {
		return
	}
}

func sampleFunc(ctx context.Context, tx *firestore.Transaction) error {
	// dummy
	var dr *firestore.DocumentRef

	if _, err := tx.Get(dr); err != nil {
		return err
	}

	if err := tx.Delete(dr); err != nil {
		return err
	}

	return nil
}
