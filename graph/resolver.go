package graph

import "helloserver/backend"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	backend *backend.Backend
}

func NewResolver(backend *backend.Backend) *Resolver {
	return &Resolver{backend}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mresolver{r.backend} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &qresolver{r.backend} }

type qresolver struct{ backend *backend.Backend }
type mresolver struct{ backend *backend.Backend }
