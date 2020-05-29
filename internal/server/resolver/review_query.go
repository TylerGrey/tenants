package resolver

import (
	"github.com/TylerGrey/tenants/internal/server/resolver/args"
	"github.com/TylerGrey/tenants/internal/server/resolver/model"
)

// Reviews ...
func (r *Resolver) Reviews(args args.ReviewsArgs) (model.ReviewConnection, error) {
	return model.ReviewConnection{}, nil
}

// Review ...
func (r *Resolver) Review(args struct {
	ID string
}) (model.Review, error) {
	return model.Review{}, nil
}
