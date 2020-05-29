package model

// ReviewConnection ...
type ReviewConnection struct {
}

// Edges ...
func (rc ReviewConnection) Edges() *[]*ReviewEdge {
	return nil
}

// Nodes ...
func (rc ReviewConnection) Nodes() *[]*Review {
	return nil
}

// PageInfo ...
func (rc ReviewConnection) PageInfo() PageInfo {
	return PageInfo{}
}

// TotalCount ...
func (rc ReviewConnection) TotalCount() int32 {
	return 0
}
