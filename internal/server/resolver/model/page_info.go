package model

// PageInfo ...
type PageInfo struct {
	Start       *string
	End         *string
	HasNext     bool
	HasPrevious bool
}

// EndCursor ...
func (pi PageInfo) EndCursor() *string {
	return pi.End
}

// HasNextPage ...
func (pi PageInfo) HasNextPage() bool {
	return pi.HasNext
}

// HasPreviousPage ...
func (pi PageInfo) HasPreviousPage() bool {
	return pi.HasPrevious
}

// StartCursor ...
func (pi PageInfo) StartCursor() *string {
	return pi.Start
}
