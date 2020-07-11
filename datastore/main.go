package main

import (
	"context"

	"cloud.google.com/go/datastore"
)

func main() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, datastore.DetectProjectID)
	if err != nil {
		return
	}

	// Pattern1: Manual NewTransaction, Commit and Rollback.
	key1 := datastore.NameKey("Dummy", "foo", nil)
	tx, err := client.NewTransaction(ctx)
	if err != nil {
		return
	}
	if err := tx.Get(key1, nil); err != nil {
		return
	}
	if _, err := tx.Put(key1, nil); err != nil {
		return
	}
	if _, err := tx.Commit(); err != nil {
		return
	}

	// Pattern2: Auto NewTransaction, Commit and Rollback
	if _, err := client.RunInTransaction(ctx, sampleFunc); err != nil {
		return
	}
}

func sampleFunc(tx *datastore.Transaction) error {
	key2 := datastore.NameKey("Dummy", "foo", nil)
	if err := tx.Get(key2, nil); err != nil {
		return err
	}
	if _, err := tx.Put(key2, nil); err != nil {
		return err
	}
	return nil
}
