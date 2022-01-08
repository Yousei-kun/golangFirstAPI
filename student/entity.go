package student

import "time"

type Student struct {
	ID        int
	SID       string
	Name      string
	Score     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
