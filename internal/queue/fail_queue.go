package queue

import (
	"context"
	"encoding/json"
	"museum_ticket/internal/jobs"

	"github.com/redis/go-redis/v9"
)

type FailQueue struct {
	client *redis.Client
}

func NewFailQueue(
	client *redis.Client,
) *FailQueue {
	return &FailQueue{
		client: client,
	}
}

func (q *FailQueue) Push(
	ctx context.Context,
	job jobs.FailJob,
) error {

	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return q.client.LPush(
		ctx,
		"failed_email_queue",
		data,
	).Err()

}
