package model

import "github.com/TylerGrey/tenants/internal/mysql/repo"

// ReviewEdge ...
type ReviewEdge struct {
	Payload repo.Review
}

// Cursor ...
func (re ReviewEdge) Cursor() string {
	return re.Payload.Cursor
}

// Node ...
func (re ReviewEdge) Node() *Review {
	return &Review{
		Payload: re.Payload,
	}
}
