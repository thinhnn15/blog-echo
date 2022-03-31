package req

type ReqNewPost struct {
	Title string `json:"title,omitempty" validate:"required"`
	Content string `json:"content,omitempty" validate:"required"`
}
