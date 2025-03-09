package worker

import (
	"context"
	"crypto/rand"
	"distribute_task_processing/config"
	"distribute_task_processing/models"
	"fmt"
	"sync"
	"time"
)

func Start(ctx context.Context, id int, taskChan <- chan models.Task, resultChan chan<- models.Result, wg *sync.WaitGroup){
	defer wg.Done()
	rateLimiter := time.Tick(config.RateLimit)

	for {
		select {
		case task, ok := <-taskChan:
			if !ok {
				fmt.Printf("Worker %d : channel closed\n",id)
				return
			}
			<-rateLimiter
			result := processTask(id,task)
			select {
			case resultChan <- result :
			case <-ctx.Done():
				return
				
			}
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down\n",id)
			return

		}
	}
}

func processTask(workerID int, task models.Task) models.Result {
	time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(200)))
	return models.Result{
		TaskID: task.ID,
		WorkerID: workerID,
		ProcessedAt: time.Now(),
		Output: fmt.Sprintf("processed-%d",task.ID),
		IsSuccess: rand.Float32()<0.8,
	}
}