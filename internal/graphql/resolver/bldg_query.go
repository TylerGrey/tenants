package resolver

import (
	"context"
	"strconv"

	"github.com/TylerGrey/tenants/internal/graphql/model"
)

// Bldgs 건물 조회
func (r *queryResolver) Bldgs(ctx context.Context, lat float64, lng float64, scale int) ([]*model.Bldg, error) {
	resolvers := []*model.Bldg{}

	// TODO scale 별 거리 조절
	bldg, err := r.BldgRepo.List(lat, lng)
	if err != nil {
		return nil, err
	}

	for _, b := range bldg {
		resolvers = append(resolvers, &model.Bldg{
			ID:          strconv.FormatUint(b.ID, 10),
			Lat:         b.Lat,
			Lng:         b.Lng,
			Rating:      b.Rating,
			Address:     b.Addr,
			RoadAddress: b.RoadAddress,
			UpdatedAt:   b.UpdatedAt.String(),
		})
	}

	return resolvers, nil
}
