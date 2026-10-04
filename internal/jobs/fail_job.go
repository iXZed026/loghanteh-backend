package jobs

import "time"

type FailJob struct {
	JobID string `json:"job_id"`

	Payload string `json:"payload"`

	Reason string `json:"reason"`

	FailedAt time.Time `json:"failed_at"`
}
