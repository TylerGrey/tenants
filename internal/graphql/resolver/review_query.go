package resolver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TylerGrey/tenants/internal/graphql/model"
	"github.com/TylerGrey/tenants/internal/mysql/repo"
)

// Reviews ...
func (r *queryResolver) Reviews(ctx context.Context, bldgID string, after *string, before *string, first *int, last *int, orderBy *model.ReviewOrder) (*model.ReviewConnection, error) {
	if first == nil && last == nil {
		return nil, fmt.Errorf("You must provide a `first` or `last` value to properly paginate the `reviews` connection")
	} else if first != nil && last != nil {
		return nil, fmt.Errorf("Passing both `first` and `last` to paginate the `reviews` connection is not supported")
	}

	id, err := strconv.ParseUint(bldgID, 10, 64)
	if err != nil {
		return nil, err
	}

	reviewsArgs := repo.ListArgs{
		First:  first,
		Last:   last,
		After:  after,
		Before: before,
	}
	if orderBy != nil {
		reviewsArgs.Order = &repo.Order{
			Field:     orderBy.Field.String(),
			Direction: orderBy.Direction.String(),
		}
	}
	reviews, page, err := r.ReviewRepo.List(id, reviewsArgs)
	if err != nil {
		return nil, err
	}

	return &model.ReviewConnection{
		Payload: reviews,
		Page:    page,
	}, nil
}

// Review ...
func (r *queryResolver) Review(ctx context.Context, id string) (*model.Review, error) {
	reviewID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	review, err := r.ReviewRepo.FindByID(reviewID)
	if err != nil {
		return nil, err
	} else if review == nil {
		return nil, nil
	}

	return nil, nil
}
