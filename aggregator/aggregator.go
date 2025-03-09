package aggregator

import (
	"context"
	"distribute_task_processing/models"
	"fmt"
	"sync"
	"sync/atomic"
)


type ResultAggregator struct {
	mu			sync.RWMutex
	results		[]models.Result
	successCnt	int32 
	errorCnt	int32 

}

func NewResultAggregator() *ResultAggregator {
	return &ResultAggregator{}
}
func ( ra *ResultAggregator) AddResult(result models.Result){
	ra.mu.Lock()
	defer ra.mu.Unlock()
	ra.results = append(ra.results, result)
	if result.IsSuccess {
		atomic.AddInt32(&ra.successCnt,1)
	}else{
		atomic.AddInt32(&ra.errorCnt,1)
	}
}


func (ra *ResultAggregator) GetResults() []models.Result {
	ra.mu.RLock()
	defer ra.mu.RUnlock()
	return ra.results
}


func (ra *ResultAggregator) Start(ctx context.Context, resultChan <- chan models.Result, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		case result:= <- resultChan:
			ra.AddResult(result)
			fmt.Printf("Processed task %d (success:%v)\n", result.TaskID, result.IsSuccess)
		case <-ctx.Done():
			fmt.Println("Aggregator stopping...")
			return 
		}
	}
}