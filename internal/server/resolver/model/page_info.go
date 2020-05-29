package model

// PageInfo ...
type PageInfo struct {
}

// EndCursor ...
func (pi PageInfo) EndCursor() *string {
	return nil
}

// HasNextPage ...
func (pi PageInfo) HasNextPage() bool {
	return false
}

// HasPreviousPage ...
func (pi PageInfo) HasPreviousPage() bool {
	return false
}

// StartCursor ...
func (pi PageInfo) StartCursor() *string {
	return nil
}
