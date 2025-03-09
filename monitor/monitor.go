package monitor

import (
	"context"
	"distribute_task_processing/config"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Monitor struct{
	taskProcessed		int32
	activeWorker		int32 

}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) Start(ctx context.Context, wg *sync.WaitGroup){
	defer wg.Done()
	ticker := time.NewTicker(config.MonitorPeriod)
	defer ticker.Stop()

	for{
		select{
		case <-ticker.C:
			fmt.Printf("\n[Monitor] Active workers:%d,Tasks processed: %d\n\n",
			atomic.LoadInt32(&m.activeWorker),
			atomic.LoadInt32(&m.taskProcessed))
		case <-ctx.Done():
			fmt.Println("Monitor stopping...")
			return
		}
	}
}