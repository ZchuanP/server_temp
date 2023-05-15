package zeebe

import (
	"context"
	"github.com/camunda/zeebe/clients/go/v8/pkg/zbc"
	"log"
	"time"
)

// 原子服务并不绑定于Process，只绑定Client。
// 仍旧可以按照现行方案进行实现。
func MustStartWorker(client zbc.Client) worker.JobWorker {
	w := client.NewJobWorker().
		JobType(kJobType).
		Handler(handleJob).
		Concurrency(1).
		MaxJobsActive(10).
		RequestTimeout(1 * time.Second).
		PollInterval(1 * time.Second).
		Name(kWorkerName).
		Open()

	log.Printf("started worker [%s] for jobs of type [%s]", kWorkerName, kJobType)
	return w
}

func handleJob(client worker.JobClient, job entities.Job) {
	worker := new(Worker)
	err := job.GetVariablesAs(worker)
	if err != nil {
		log.Printf("failed to get variables for job %d: [%s]", job.Key, err)
		return
	}

	code, err := worker.Do(context.Background())
	if err != nil {
		client.NewThrowErrorCommand().JobKey(job.Key).ErrorCode(code)
		return
	}

	_, err = client.NewCompleteJobCommand().JobKey(job.Key).VariablesFromObject(worker)
	if err != nil {
		log.Printf("failed to complete job with key %d: [%s]", job.Key, err)
		return
	}

	log.Printf("completed job %d successfully", job.Key)
}
