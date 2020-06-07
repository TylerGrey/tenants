package resolver

import "github.com/TylerGrey/tenants/internal/server/resolver/model"

// Bldgs 건물 조회
func (r Resolver) Bldgs(args struct {
	Lat   float64
	Lng   float64
	Scale int32
}) ([]model.Bldg, error) {
	resolvers := []model.Bldg{}

	// TODO scale 별 거리 조절
	bldg, err := r.BldgRepo.List(args.Lat, args.Lng)
	if err != nil {
		return resolvers, err
	}

	for _, b := range bldg {
		resolvers = append(resolvers, model.Bldg{
			Payload: *b,
		})
	}

	return resolvers, nil
}
