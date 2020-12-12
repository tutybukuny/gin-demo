package serializers

type BookSerializer struct {
	Title       string  `json:"title" binding:"required"`
	Author      string  `json:"author" binding:"required"`
	Description string  `json:"description"`
	Count       float32 `json:"count"`
}
