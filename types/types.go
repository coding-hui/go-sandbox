package types

import "time"

// Result is the result returned by strat trace
type Result struct {
	Status            // the final status for the process
	ExitStatus int    // exit Status
	UserTime   uint64 // used user CPU time (in ms)
	UserMem    uint64 // used user memory (in kb)
	Stat              // collects time usage for the runner
}

// Stat is the time usages in ns
type Stat struct {
	SetUpTime   time.Duration
	RunningTime time.Duration
}

// Limit represents the resource limit for traced process
type Limit struct {
	TimeLimit   uint64 // user CPU time limit (in ms)
	MemoryLimit uint64 // user memory limit (in kB)
}
