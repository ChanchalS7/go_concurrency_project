package main

import (
	"context"
	"distribute_task_processing/aggregator"
	"distribute_task_processing/config"
	"distribute_task_processing/models"
	"distribute_task_processing/monitor"
	"distribute_task_processing/producer"
	"distribute_task_processing/server"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	//initialize components
	taskChan :=make(chan models.Task, config.QueueSize)
	resultChan := make(chan models.Result, config.QueueSize)
	aggregator := aggregator.NewResultAggregator()
	monitor := monitor.NewMonitor()

	//start producer 
	wg.Add(1)
	go producer.Start(ctx, taskChan, wg)

	// Start aggregator
	wg.Add(1)
	go aggregator.Start(ctx, resultChan, wg)

	// Start monitor
	wg.Add(1)
	go monitor.Start(ctx, wg)

	// Start HTTP server
	wg.Add(1)
	go server.Start(ctx, aggregator, wg)

	// Handle shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	<-sigChan
	fmt.Println("\nReceived interrupt - shutting down...")
	cancel()
	wg.Wait()
	fmt.Println("Shutdown complete")

}