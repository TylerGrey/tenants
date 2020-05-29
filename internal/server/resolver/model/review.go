package model

import (
	"strconv"

	"github.com/TylerGrey/tenants/internal/mysql/repo"
	"github.com/graph-gophers/graphql-go"
)

// Review ...
type Review struct {
	Payload repo.Review
}

// ID ...
func (r Review) ID() graphql.ID {
	id := strconv.FormatUint(r.Payload.ID, 10)
	return graphql.ID(id)
}

// Title ...
func (r Review) Title() string {
	return r.Payload.Title
}

// Content ...
func (r Review) Content() string {
	return r.Payload.Content
}

// TotalScore ...
func (r Review) TotalScore() float64 {
	sum := float64(r.Payload.ScoreRent + r.Payload.ScoreConvenience + r.Payload.ScoreLandlord + r.Payload.ScoreMaintenanceFees + r.Payload.ScorePublicTransport)
	return sum / 5
}

// Score ...
func (r Review) Score() ReviewScore {
	return ReviewScore{
		ScoreRent:            r.Payload.ScoreRent,
		ScoreMaintenanceFees: r.Payload.ScoreMaintenanceFees,
		ScorePublicTransport: r.Payload.ScorePublicTransport,
		ScoreConvenience:     r.Payload.ScoreConvenience,
		ScoreLandlord:        r.Payload.ScoreLandlord,
	}
}

// UpdatedAt ...
func (r Review) UpdatedAt() string {
	return r.Payload.UpdatedAt.String()
}

// ReviewScore ...
type ReviewScore struct {
	ScoreRent            int32
	ScoreMaintenanceFees int32
	ScorePublicTransport int32
	ScoreConvenience     int32
	ScoreLandlord        int32
}

// Rent ...
func (rc ReviewScore) Rent() int32 {
	return rc.ScoreRent
}

// MaintenanceFees ...
func (rc ReviewScore) MaintenanceFees() int32 {
	return rc.ScoreMaintenanceFees
}

// PublicTransport ...
func (rc ReviewScore) PublicTransport() int32 {
	return rc.ScorePublicTransport
}

// Convenience ...
func (rc ReviewScore) Convenience() int32 {
	return rc.ScoreConvenience
}

// Landlord ...
func (rc ReviewScore) Landlord() int32 {
	return rc.ScoreLandlord
}
