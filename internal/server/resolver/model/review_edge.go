package model

// ReviewEdge ...
type ReviewEdge struct {
}

// Cursor ...
func (re ReviewEdge) Cursor() string {
	return "cursor"
}

// Node ...
func (re ReviewEdge) Node() *Review {
	return nil
}
