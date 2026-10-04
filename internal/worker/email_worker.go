package worker

import (
	"context"
	"encoding/json"
	"log/slog"
	"museum_ticket/internal/jobs"
	"museum_ticket/internal/mail"
	"museum_ticket/internal/queue"
	"strconv"
	"time"

	"github.com/bytedance/gopkg/util/logger"
)

type EmailWorker struct {
	queue     *queue.RedisQueue
	failQueue *queue.FailQueue
	mailer    mail.Mailer
	logger    *slog.Logger

	emailWorkerCount string
}

func NewEmailWorker(
	queue *queue.RedisQueue,
	failQueue *queue.FailQueue,
	mailer mail.Mailer,
	logger *slog.Logger,
	emailWorker string,
) *EmailWorker {

	return &EmailWorker{
		queue:            queue,
		failQueue:        failQueue,
		mailer:           mailer,
		logger:           logger,
		emailWorkerCount: emailWorker,
	}

}

const maxRetry = 3

func (w *EmailWorker) Start(
	ctx context.Context,
) {

	workerCount, err := strconv.Atoi(w.emailWorkerCount)
	if err != nil {
		logger.Error(
			"converting the string to a number resulted in an error.",
			"error", err,
		)
		return
	}

	for i := 1; i <= workerCount; i++ {

		go w.workerLoop(
			ctx,
			i,
		)

	}

	<-ctx.Done()

}

func (w *EmailWorker) workerLoop(

	ctx context.Context,

	id int,

) {

	for {

		select {

		case <-ctx.Done():

			w.logger.Info(
				"worker stopped",
				"worker",
				id,
			)

			return

		default:

			job, err := w.queue.PopEmailJob(ctx)

			if err != nil {

				w.logger.Error(
					"get job failed",
					"worker",
					id,
					"error",
					err,
				)

				continue
			}

			w.process(
				ctx,
				job,
			)

		}

	}

}

func (w *EmailWorker) process(
	ctx context.Context,
	job jobs.EmailJob,
) {

	err :=
		w.mailer.SendTicketEmail(
			ctx,
			job.Email,
			job.TicketID,
		)

	if err != nil {

		job.Attempt++

		if job.Attempt < maxRetry {

			w.logger.Info(
				"email failed, retrying",
				"job_id",
				job.ID,
				"attempts",
				job.Attempt,
			)

			if err := w.queue.PushEmailJob(ctx, job); err != nil {
				w.logger.Error(
					"failed queue push error",
					"error",
					err,
				)
			}
			return
		}
		data, _ := json.Marshal(job)

		failJob := jobs.FailJob{
			JobID:    job.ID,
			Payload:  string(data),
			Reason:   err.Error(),
			FailedAt: time.Now(),
		}

		if err := w.failQueue.Push(ctx, failJob); err != nil {
			w.logger.Error(
				"failed queue push error",
				"error",
				err,
			)

			return
		}

		return

	}

	w.logger.Info(
		"email sent",
		"job_id",
		job.ID,
	)

}
