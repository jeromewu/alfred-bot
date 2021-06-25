package models

import (
	"context"

	"cloud.google.com/go/firestore"
)

func ReadTODOs(conf *Conf, userId string) (*Todo, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, conf.ProjectID)
	defer client.Close()
	if err != nil {
		return nil, err
	}

	todoDoc := client.Doc("todos/" + userId)
	todo := new(Todo)

	snap, err := todoDoc.Get(ctx)
	if err != nil {
		_, err = todoDoc.Set(ctx, todo)
		if err != nil {
			return nil, err
		}
		snap, _ = todoDoc.Get(ctx)
	}

	if err := snap.DataTo(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func WriteTODOs(conf *Conf, userId string, todo *Todo) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, conf.ProjectID)
	defer client.Close()
	if err != nil {
		return err
	}

	todoDoc := client.Doc("todos/" + userId)

	_, err = todoDoc.Set(ctx, todo)
	if err != nil {
		return err
	}

	return nil
}
