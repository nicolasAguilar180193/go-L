package domain

import "time"

// Player represents a player in the domain.
type Player struct {
	ID          any          `json:"id" bson:"_id"`
	FirstName   string       `json:"first_name" bson:"first_name"`
	LastName    string       `json:"last_name" bson:"last_name"`
	ContactInfo *ContactInfo `json:"contact_info" bson:"contact_info"`
	DateOfBirth *Date        `json:"date_of_birth"`
	TeamInfo    *TeamInfo    `json:"team_info" bson:"team_info"`
	CreatedAt   *time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at" bson:"updated_at"`
}
