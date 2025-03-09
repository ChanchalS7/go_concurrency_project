package main

import "time"

//Task represents a unit fo work
type Task struct {
	ID				int 
	Data			string
	CreatedAt		time.Time	
}

//Result contains task processing results 
type Result struct {
	TaskID			int 
	WorkerID		int 
	ProcessedAt		time.Time
	Output			string 
	IsSuccess		bool 
}