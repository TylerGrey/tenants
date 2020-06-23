package model

import "github.com/TylerGrey/tenants/internal/mysql/repo"

// ReviewConnection ...
type ReviewConnection struct {
	Payload []*repo.Review
	Page    repo.PageInfo
}

// Edges ...
func (rc ReviewConnection) Edges() []*ReviewEdge {
	edges := []*ReviewEdge{}

	for _, p := range rc.Payload {
		edges = append(edges, &ReviewEdge{
			Payload: *p,
		})
	}

	return edges
}

// Nodes ...
func (rc ReviewConnection) Nodes() []*Review {
	nodes := []*Review{}

	for _, d := range rc.Payload {
		nodes = append(nodes, &Review{
			Payload: *d,
		})
	}

	return nodes
}

// PageInfo ...
func (rc ReviewConnection) PageInfo() PageInfo {
	pageInfo := PageInfo{
		HasNextPage:     rc.Page.HasNext,
		HasPreviousPage: rc.Page.HasPrev,
	}

	if len(rc.Payload) > 0 {
		pageInfo.StartCursor = &rc.Payload[0].Cursor
		pageInfo.EndCursor = &rc.Payload[len(rc.Payload)-1].Cursor
	}

	return pageInfo
}

// TotalCount ...
func (rc ReviewConnection) TotalCount() int32 {
	return rc.Page.Total
}
