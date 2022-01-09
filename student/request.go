package student

type StudentRequest struct {
	SID   string      `json:"sid" binding:"required"`
	Name  string      `json:"name" binding:"required"`
	Score interface{} `json:"score" binding:"required,number"`
}
