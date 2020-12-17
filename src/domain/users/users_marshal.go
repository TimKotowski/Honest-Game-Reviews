package users

type PublicUser struct {
	Username string `json:"user_name"`
	Email    string `json:"email"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	Username    string `json:"user_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) MarshalUser(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Username: user.Username,
			Email:    user.Email,
		}
	}
	return PrivateUser{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		DateCreated: user.DateCreated,
	}
}
