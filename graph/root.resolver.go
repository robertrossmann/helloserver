package graph

import (
	"context"
	"helloserver/graph/gqlm"
	"time"
)

func (r *qresolver) Root(_ context.Context) (*gqlm.Root, error) {
	root := &gqlm.Root{
		Now: time.Now().Format(time.RFC3339),
		Env: "local",
	}

	return root, nil
}

func (r *mresolver) Empty(_ context.Context) (*bool, error) {
	empty := true
	return &empty, nil
}
