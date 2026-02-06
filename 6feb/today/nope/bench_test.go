package main

import (
	"fmt"
	"testing"
)

func BenchmarkWorkerPool_SmallLoad(b *testing.B) {
	numJobs := 100
	numWorkers := 4

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RunWorkerPool(numJobs, numWorkers)
	}
}

func BenchmarkWorkerPool_MediumLoad(b *testing.B) {
	numJobs := 1000
	numWorkers := 40 // make both 10 times

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RunWorkerPool(numJobs, numWorkers)
	}
}

func BenchmarkWorkerPool_HighLoad(b *testing.B) {
	numJobs := 5000
	numWorkers := 200 // make both 5 times from medium

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RunWorkerPool(numJobs, numWorkers)
	}
}

func BenchmarkWorkerPool_WorkerScaling(b *testing.B) {

	numJobs := 5000 // fixed workload

	workerCases := []int{
		1,
		2,
		4,
		8,
		16,
		32,
		64,
		128,
		264,
	}

	for _, workers := range workerCases {

		workers := workers // capture loop variable

		b.Run(fmt.Sprintf("Workers_%d", workers), func(b *testing.B) {

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				RunWorkerPool(numJobs, workers)
			}
		})
	}
}

func BenchmarkWorkerPool_jobScaling(b *testing.B) {

	workerCases := 50// fixed worker

	numJobs := []int{
		500,
		1000,
		2000,
		4000,
		8000,
		16000,
		32000,
		64000,
		128000,
		256000,
	}

	for _, job := range numJobs {

		job := job // capture loop variable

		b.Run(fmt.Sprintf("Workers_%d", job), func(b *testing.B) {

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				RunWorkerPool(job, workerCases)
			}
		})
	}
}