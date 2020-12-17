package reviews_handler

import (
	"context"
	"net/http"
)

type ReviewsHandlerInterface interface {
	GetReviews(w http.ResponseWriter, r *http.Request)
	UpdateReview(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type reviewsHandler struct{}

func NewReviewsHandler() ReviewsHandlerInterface {
	return &reviewsHandler{}
}

func (handler *reviewsHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	// write a service to get back a reivew
	// service should have validation
}

func (handler *reviewsHandler) UpdateReview(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// middleware to confirm that this user can update the review if its there review
	// write service to confirm the review was update with the update reivew
	// provide a PAtch for a partial update

}
func (handler *reviewsHandler) DeleteReview(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// middleware to confirm that this user can delete the review if its there review
	// write service to confirm the review was deleted with the update reivew
}
