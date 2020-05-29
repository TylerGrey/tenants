package resolver

import (
	"strconv"

	"github.com/TylerGrey/tenants/internal/mysql/repo"
	"github.com/TylerGrey/tenants/internal/server/resolver/args"
	"github.com/TylerGrey/tenants/internal/server/resolver/model"
)

// CreateReview ...
func (r *Resolver) CreateReview(args args.CreateReviewArgs) (*model.Review, error) {
	bldg, err := r.BldgRepo.FindByLatLng(args.Input.Lat, args.Input.Lng)
	if err != nil {
		return nil, err
	}

	if bldg == nil {
		if bldg, err = r.BldgRepo.Create(repo.Bldg{
			Lat: args.Input.Lat,
			Lng: args.Input.Lng,
		}); err != nil {
			return nil, err
		}
	}

	m := repo.Review{
		BldgID:               bldg.ID,
		Title:                args.Input.Title,
		Content:              args.Input.Content,
		ScoreRent:            args.Input.Score.Rent,
		ScoreConvenience:     args.Input.Score.Convenience,
		ScoreLandlord:        args.Input.Score.Landlord,
		ScoreMaintenanceFees: args.Input.Score.MaintenanceFees,
		ScorePublicTransport: args.Input.Score.PublicTransport,
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
func (r *Resolver) UpdateReview(args args.UpdateReviewArgs) (*model.Review, error) {
	id, err := strconv.ParseUint(args.Input.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	m := repo.Review{
		ID: id,
	}

	if args.Input.Title != nil {
		m.Title = *args.Input.Title
	}

	if args.Input.Content != nil {
		m.Content = *args.Input.Content
	}

	if args.Input.Score != nil {
		m.ScoreRent = args.Input.Score.Rent
		m.ScoreConvenience = args.Input.Score.Convenience
		m.ScoreLandlord = args.Input.Score.Landlord
		m.ScoreMaintenanceFees = args.Input.Score.MaintenanceFees
		m.ScorePublicTransport = args.Input.Score.PublicTransport
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
func (r *Resolver) DeleteReview(args struct {
	ID string
}) (bool, error) {
	id, err := strconv.ParseUint(args.ID, 10, 64)
	if err != nil {
		return false, err
	}

	return r.ReviewRepo.Delete(id)
}
