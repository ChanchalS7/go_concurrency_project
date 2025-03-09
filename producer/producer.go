package producer

import (
	"context"
	"crypto/rand"
	"distribute_task_processing/models"
	"fmt"
	"sync"
	"time"
)

func Start( ctx context.Context, taskChan chan <- models.Task, wg *sync.WaitGroup){
	defer wg.Done()
	defer close(taskChan)

	taskID := 1

	for {
		select {
		case <- ctx.Done():
				fmt.Println("Producer stopping...")
				return 
		default:
			task :=models.Task{
				ID:		taskID,
				Data:   fmt.Sprintf("data-%d", taskID),
				CreatedAt: time.Now(),	
			}
			select {
			case taskChan <- task:
				fmt.Printf("Produced task %d\n",taskID)
				taskID++
			case <- ctx.Done():
				return	
			}
			time.Sleep(time.Millisecond * time.Duration(100+rand.Intn(400)))
					
		}
	}
}