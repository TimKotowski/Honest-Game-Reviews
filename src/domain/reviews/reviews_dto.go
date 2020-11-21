package reviews

type Reviews struct {
	ID        int    `json:"id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	Recommend bool   `json:"recommend"`
	UserID    int64  `json:"user_id"`
	GameID    int64  `json:"game_id"`
}
