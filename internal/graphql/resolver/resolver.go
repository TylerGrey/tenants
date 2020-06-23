package resolver

import (
	"github.com/TylerGrey/tenants/internal/graphql/generated"
	"github.com/TylerGrey/tenants/internal/mysql/repo"
)

// Resolver ...
type Resolver struct {
	ReviewRepo repo.ReviewRepository
	BldgRepo   repo.BldgRepository
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
