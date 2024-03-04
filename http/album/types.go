package album

type Album struct {
	ID     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}
