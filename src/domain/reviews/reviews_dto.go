package reviews

import (
	"Honest-Game-Reviews/src/utils/errors"
)

type Review struct {
	ID int64 `json:"id"`
	// low optimization but rating wont exceed 5
	Rating  int8   `json:"rating"`
	Comment string `json:"comment"`
	GameID  int    `json:"game_id"`
	UserID  int    `json:"user_id"`
}

func (review *Review) Validate() *errors.RestErrors {
	if review.Rating == 0 {
		return errors.NewBadRequestError("no rating given")
	}
	if review.Comment == "" || len(review.Comment) <= 10 {
		return errors.NewBadRequestError("no comment given, or provide a longer comment to be a valid comment")
	}
	return nil
}
