package server

import (
	"context"
	"distribute_task_processing/aggregator"
	"distribute_task_processing/config"

	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Start(ctx context.Context, aggregator *aggregator.ResultAggregator, wg *sync.WaitGroup){
	defer wg.Done()
	server  := &http.Server{Addr: config.HTTPPort}
	http.HandleFunc("/results",func(w http.ResponseWriter, r *http.Request){
		results:= aggregator.GetResults()
		json.NewEncoder(w).Encode(results)
	})
	go func ()  {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %v\n",err)
		}

	}()
	<- ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("HTTP shutdown error: %v\n",err)
	}
}