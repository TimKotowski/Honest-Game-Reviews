package reviews_service

var (
	ReviewsServices ReviewsService = &reviewsService{}
)

type ReviewsService interface {
	GetReviews()
	GetReviewByID()
	UpdateReview()
	DeleteReview()
}

type reviewsService struct{}

func (s reviewsService) GetReviews() {}

func (s reviewsService) GetReviewByID() {}

func (s reviewsService) UpdateReview() {}

func (s reviewsService) DeleteReview() {}
