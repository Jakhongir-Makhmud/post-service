package structs

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostParams struct {
	Page  int `json:"offset"`
	Limit int `json:"limit"`
}
