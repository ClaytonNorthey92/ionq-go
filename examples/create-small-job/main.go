package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	ionq "github.com/ClaytonNorthey92/ionq-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	client := ionq.NewClient(
		"https://api.ionq.co/v0.3",
		os.Getenv("IONQ_API_KEY"),
	)

	target1 := uint(0)
	target2 := uint(1)

	response, err := client.CreateJob(ctx, &ionq.CreateJobRequest{
		Input: ionq.JobInput{
			Format: "ionq.circuit.v0",
			Qubits: 2,
			Circuit: []ionq.CircuitInput{
				{
					Gate:   "h",
					Target: &target1,
				},
				{
					Gate:   "h",
					Target: &target2,
				},
			},
		},
	})

	fmt.Printf("received response: %v\n", response)

	if err != nil {
		panic(fmt.Sprintf("error creating job: %s", err))
	}

	if response.Status != http.StatusOK {
		panic(fmt.Sprintf("received unexpected http status code: %d", response.Status))
	}

	jobId := response.Response.ID

	for {
		select {
		case <-ctx.Done():
			panic(ctx.Err().Error())
		case <-time.After(5 * time.Second):
			response, err := client.GetJob(ctx, &ionq.GetJobRequest{
				ID: jobId,
			})
			if err != nil {
				panic(fmt.Sprintf("error getting job: %s", err))
			}

			if response.Status != http.StatusOK {
				panic(fmt.Sprintf("received unexpected http status code: %d", response.Status))
			}

			fmt.Printf("received job status of %s\n", response.Response.Status)

			if response.Response.Status == "completed" {
				break
			}
		}
		break
	}

	outputResponse, err := client.GetJobOutput(ctx, &ionq.GetJobOutputRequest{
		ID: jobId,
	})
	if err != nil {
		panic(fmt.Sprintf("error getting job: %s", err))
	}

	if response.Status != http.StatusOK {
		panic(fmt.Sprintf("received unexpected http status code: %d", response.Status))
	}

	fmt.Printf("job output is %v\n", outputResponse.Response)
}
