package queue

import (
	"context"
	"encoding/json"
	"museum_ticket/internal/jobs"
)

func (q *RedisQueue) PushEmailJob(
	ctx context.Context,
	job jobs.EmailJob,
) error {

	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return q.client.LPush(
		ctx,
		"email_queue",
		data,
	).Err()

}

func (q *RedisQueue) PopEmailJob(
	ctx context.Context,
) (jobs.EmailJob, error) {

	result, err :=
		q.client.BRPop(
			ctx,
			0,
			"email_queue",
		).Result()

	if err != nil {

		return jobs.EmailJob{}, err
	}

	var job jobs.EmailJob

	if err :=
		json.Unmarshal(
			[]byte(result[1]),
			&job,
		); err != nil {
		return jobs.EmailJob{}, err
	}

	return job, nil

}
