package domain

import "time"

type Player struct {
	Name      string `json:"name" binding:"required"`
	Age       int    `json:"age" binding:"required,min=0"`
	CreatedAt time.Time
}
