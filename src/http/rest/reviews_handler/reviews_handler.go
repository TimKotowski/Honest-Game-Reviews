package reviews_handler

import (
	"Honest-Game-Reviews/src/utils/json_utils"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ReviewsHandlerInterface interface {
	GetReviews(w http.ResponseWriter, r *http.Request)
	GetReviewByID(w http.ResponseWriter, r *http.Request)
	UpdateReview(ctx context.Context, w http.ResponseWriter, r *http.Request)
	DeleteReview(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type reviewsHandler struct{}

// this function is meant to be used to be called in the api routes in urlReviewMapping file
func NewReviewsHandler() ReviewsHandlerInterface {
	return &reviewsHandler{}
}

func (handler *reviewsHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	// write a service to get back a reivew
	// service should have validation

}

func (handler *reviewsHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	id, parseIntErr := strconv.ParseInt(chi.URLParam(r, "review_id"), 10, 64)
	if parseIntErr != nil {
		json_utils.ClientErrorResponse(w, http.StatusBadRequest, "invalid param value", parseIntErr)
	}
	fmt.Println(id)
	// service
	// reviews_service.ReviewsServices.GetReviewByID()
	// write a service to get back review based on id paramater send from client
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
