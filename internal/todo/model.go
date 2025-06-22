package todo

type Todo struct {
	ID    int    `json:"id" example:"1"`
	Title string `json:"title" example:"Learn Go"`
	Done  bool   `json:"done" example:"false"`
}
