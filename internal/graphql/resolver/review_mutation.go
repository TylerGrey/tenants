package resolver

import (
	"context"
	"strconv"

	"github.com/TylerGrey/tenants/internal/graphql/model"
	"github.com/TylerGrey/tenants/internal/mysql/repo"
)

// CreateReview ...
func (r *mutationResolver) CreateReview(ctx context.Context, input model.CreateReviewInput) (*model.Review, error) {
	bldg, err := r.BldgRepo.FindByLatLng(input.Lat, input.Lng)
	if err != nil {
		return nil, err
	}

	if bldg == nil {
		if bldg, err = r.BldgRepo.Create(repo.Bldg{
			Lat: input.Lat,
			Lng: input.Lng,
		}); err != nil {
			return nil, err
		}
	}

	m := repo.Review{
		BldgID:               bldg.ID,
		Title:                input.Title,
		Content:              input.Content,
		ScoreRent:            int32(input.Score.Rent),
		ScoreConvenience:     int32(input.Score.Convenience),
		ScoreLandlord:        int32(input.Score.Landlord),
		ScoreMaintenanceFees: int32(input.Score.MaintenanceFees),
		ScorePublicTransport: int32(input.Score.PublicTransport),
	}
	review, err := r.ReviewRepo.Create(m)
	if err != nil {
		return nil, err
	}

	return &model.Review{
		Payload: *review,
	}, nil
}

// UpdateReview ...
func (r *mutationResolver) UpdateReview(ctx context.Context, input model.UpdateReviewInput) (*model.Review, error) {
	id, err := strconv.ParseUint(input.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	m := repo.Review{
		ID: id,
	}

	if input.Title != nil {
		m.Title = *input.Title
	}

	if input.Content != nil {
		m.Content = *input.Content
	}

	if input.Score != nil {
		m.ScoreRent = int32(input.Score.Rent)
		m.ScoreConvenience = int32(input.Score.Convenience)
		m.ScoreLandlord = int32(input.Score.Landlord)
		m.ScoreMaintenanceFees = int32(input.Score.MaintenanceFees)
		m.ScorePublicTransport = int32(input.Score.PublicTransport)
	}

	review, err := r.ReviewRepo.Update(m)
	if err != nil {
		return nil, err
	}

	return &model.Review{
		Payload: *review,
	}, nil
}

// DeleteReview ...
func (r *mutationResolver) DeleteReview(ctx context.Context, id string) (bool, error) {
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return false, err
	}

	return r.ReviewRepo.Delete(uintID)
}
