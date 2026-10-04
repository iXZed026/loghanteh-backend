package jobs

import "time"

type EmailJob struct {
	ID        string    `json:"id"`
	TicketID  string    `json:"ticket_id"`
	Email     string    `json:"email"`
	Attempt   int       `json:"attempt"`
	CreatedAt time.Time `json:"created_at"`
}
