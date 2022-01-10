package student

type StudentResponse struct {
	ID    int         `json:"id"`
	SID   string      `json:"sid"`
	Name  string      `json:"name"`
	Score interface{} `json:"score"`
}
