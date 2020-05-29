package args

// ReviewOrderBy ...
type ReviewOrderBy struct {
	Field     string
	Direction string
}

// ReviewsArgs ...
type ReviewsArgs struct {
	BldgID  string
	After   *string
	Before  *string
	First   *int32
	Last    *int32
	OrderBy *ReviewOrderBy
}

// CreateReviewArgs ...
type CreateReviewArgs struct {
	Input CreateReviewInput
}

// CreateReviewInput ...
type CreateReviewInput struct {
	Lat     float64
	Lng     float64
	Title   string
	Content string
	Score   ReviewScoreInput
}

// UpdateReviewArgs ...
type UpdateReviewArgs struct {
	Input UpdateReviewInput
}

// UpdateReviewInput ...
type UpdateReviewInput struct {
	ID      string
	Title   *string
	Content *string
	Score   *ReviewScoreInput
}

// ReviewScoreInput ...
type ReviewScoreInput struct {
	Rent            int32
	MaintenanceFees int32
	PublicTransport int32
	Convenience     int32
	Landlord        int32
}
