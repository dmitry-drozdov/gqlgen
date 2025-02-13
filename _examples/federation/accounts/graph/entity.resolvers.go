package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/dmitry-drozdov/gqlgen version v0.17.64-dev

import (
	"context"

	"github.com/dmitry-drozdov/gqlgen/_examples/federation/accounts/graph/model"
)

// FindEmailHostByID is the resolver for the findEmailHostByID field.
func (r *entityResolver) FindEmailHostByID(ctx context.Context, id string) (*model.EmailHost, error) {
	return r.HostForUserID(id)
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	name := "User " + id
	if id == "1234" {
		name = "Me"
	}

	return &model.User{
		ID:       id,
		Username: name,
		Email:    id + "@test.com",
	}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
