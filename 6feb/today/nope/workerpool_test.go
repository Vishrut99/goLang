package main

import (
	"fmt"
	"testing"
)

func TestWorkerPoolBasic(t *testing.T) {

	type testCase struct {
		numJobs    int
		numWorkers int
	}

	testCases := []testCase{
		{numJobs: 5, numWorkers: 3},
		{numJobs: 10, numWorkers: 2},
		{numJobs: 0, numWorkers: 5},
	}

	for _, tc := range testCases {

		tc := tc // IMPORTANT: capture loop variable

		t.Run(
			fmt.Sprintf("Jobs=%d_Workers=%d", tc.numJobs, tc.numWorkers),
			func(t *testing.T) {

				out := RunWorkerPool(tc.numJobs, tc.numWorkers)

				// ✅ Count check
				if len(out) != tc.numJobs {
					t.Fatalf("Expected %d results, got %d", tc.numJobs, len(out))
				}

				// ✅ Value correctness check (order independent)
				expectedMap := make(map[int]int)
				for j := 1; j <= tc.numJobs; j++ {
					expectedMap[j]++
				}

				for _, val := range out {
					if expectedMap[val] == 0 {
						t.Fatalf("Unexpected result value: %d", val)
					}
					expectedMap[val]--
				}
			},
		)
	}
}

