package resolver

import "context"

// Ping ...
func (r *Resolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}
