package resolver

import (
	"fmt"
	"strconv"

	"github.com/TylerGrey/tenants/internal/mysql/repo"
	"github.com/TylerGrey/tenants/internal/server/resolver/args"
	"github.com/TylerGrey/tenants/internal/server/resolver/model"
)

// Reviews ...
func (r *Resolver) Reviews(args args.ReviewsArgs) (model.ReviewConnection, error) {
	if args.First == nil && args.Last == nil {
		return model.ReviewConnection{}, fmt.Errorf("You must provide a `first` or `last` value to properly paginate the `reviews` connection")
	} else if args.First != nil && args.Last != nil {
		return model.ReviewConnection{}, fmt.Errorf("Passing both `first` and `last` to paginate the `reviews` connection is not supported")
	}

	bldgID, err := strconv.ParseUint(args.BldgID, 10, 64)
	if err != nil {
		return model.ReviewConnection{}, err
	}

	reviewsArgs := repo.ListArgs{
		First:  args.First,
		Last:   args.Last,
		After:  args.After,
		Before: args.Before,
	}
	if args.OrderBy != nil {
		reviewsArgs.Order = &repo.Order{
			Field:     args.OrderBy.Field,
			Direction: args.OrderBy.Direction,
		}
	}
	reviews, page, err := r.ReviewRepo.List(bldgID, reviewsArgs)
	if err != nil {
		return model.ReviewConnection{}, err
	}

	return model.ReviewConnection{
		Payload: reviews,
		Page:    page,
	}, nil
}

// Review ...
func (r *Resolver) Review(args struct {
	ID string
}) (*model.Review, error) {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	review, err := r.ReviewRepo.FindByID(id)
	if err != nil {
		return nil, err
	} else if review == nil {
		return nil, nil
	}

	return &model.Review{
		Payload: *review,
	}, nil
}
