package model

import (
	"strconv"

	"github.com/TylerGrey/tenants/internal/mysql/repo"
	"github.com/graph-gophers/graphql-go"
)

// Bldg ...
type Bldg struct {
	Payload repo.Bldg
}

// ID ...
func (b Bldg) ID() graphql.ID {
	id := strconv.FormatUint(b.Payload.ID, 10)
	return graphql.ID(id)
}

// Lat ...
func (b Bldg) Lat() float64 {
	return b.Payload.Lat
}

// Lng ...
func (b Bldg) Lng() float64 {
	return b.Payload.Lng
}

// Rating ...
func (b Bldg) Rating() float64 {
	return b.Payload.Rating
}

// UpdatedAt ...
func (b Bldg) UpdatedAt() string {
	return b.Payload.UpdatedAt.String()
}
