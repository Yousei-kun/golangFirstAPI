package student

type StudentInput struct {
	ID    string      `json:"id" binding:"required"`
	Name  string      `json:"name" binding:"required"`
	Score interface{} `json:"score" binding:"required,number"`
}
