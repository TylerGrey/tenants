package resolver

import "github.com/TylerGrey/tenants/internal/mysql/repo"

// Resolver ...
type Resolver struct {
	ReviewRepo repo.ReviewRepository
	BldgRepo   repo.BldgRepository
}
