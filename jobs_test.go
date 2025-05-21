package ionq

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-test/deep"
	"github.com/h2non/gock"
)

const (
	myFakeEndpoint = "https://myfakeionq.test/v0.3"
	myFakeAPIKey   = "blahblahnotreal"
)

func TestGetJobsSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	var jobsResponseMock GetJobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(jobsPath).
		Reply(200).
		JSON(&jobsResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobsResponseWithStatus, err := client.GetJobs(ctx, &GetJobsRequest{
		IDs: []string{"cb6d30f7-63c2-4860-9f0e-ad15cd4e2379", "e759e916-af08-4716-9b3d-15bd1bf65ffe"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobsResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobsResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobsResponseMock, jobsResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestGetJobsSuccessWithQueryParams(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var jobsResponseMock GetJobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(jobsPath).
		MatchParams(map[string]string{
			"id":     "cb6d30f7-63c2-4860-9f0e-ad15cd4e2379",
			"status": "completed",
			"limit":  "4",
			"next":   "f759e916-af08-4716-9b3d-15bd1bf65ffe",
		}).
		Reply(200).
		JSON(&jobsResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobsResponseWithStatus, err := client.GetJobs(ctx, &GetJobsRequest{
		IDs:    []string{"cb6d30f7-63c2-4860-9f0e-ad15cd4e2379", "e759e916-af08-4716-9b3d-15bd1bf65ffe"},
		Status: "completed",
		Limit:  4,
		Next:   "f759e916-af08-4716-9b3d-15bd1bf65ffe",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobsResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobsResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobsResponseMock, jobsResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestGetJobsSuccessWithQueryParamsOtherId(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var jobsResponseMock GetJobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(jobsPath).
		MatchParams(map[string]string{
			"id":     "e759e916-af08-4716-9b3d-15bd1bf65ffe",
			"status": "completed",
			"limit":  "4",
			"next":   "f759e916-af08-4716-9b3d-15bd1bf65ffe",
		}).
		Reply(200).
		JSON(&jobsResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobsResponseWithStatus, err := client.GetJobs(ctx, &GetJobsRequest{
		IDs:    []string{"cb6d30f7-63c2-4860-9f0e-ad15cd4e2379", "e759e916-af08-4716-9b3d-15bd1bf65ffe"},
		Status: "completed",
		Limit:  4,
		Next:   "f759e916-af08-4716-9b3d-15bd1bf65ffe",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobsResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobsResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobsResponseMock, jobsResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestGetJobsErrorStatusCode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var jobsResponseMock GetJobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(jobsPath).
		Reply(400).
		JSON(&jobsResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobsResponseWithStatus, err := client.GetJobs(ctx, &GetJobsRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if jobsResponseWithStatus.Status != http.StatusBadRequest {
		t.Fatalf("unexpected status: %d", jobsResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobsResponseMock, jobsResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestCreateJobsSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var createJobResponse CreateJobResponse
	if err := gofakeit.Struct(&createJobResponse); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&createJobResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Post(jobsPath).
		Reply(200).
		JSON(&createJobResponse)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	createJobWithStatus, err := client.CreateJob(ctx, &CreateJobRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if createJobWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", createJobWithStatus.Status)
	}

	if diff := deep.Equal(createJobResponse, createJobWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestCreateJobsFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var createJobResponse CreateJobResponse
	if err := gofakeit.Struct(&createJobResponse); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&createJobResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Post(jobsPath).
		Reply(400).
		JSON(&createJobResponse)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	createJobResponseWithStatus, err := client.CreateJob(ctx, &CreateJobRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if createJobResponseWithStatus.Status != 400 {
		t.Fatalf("unexpected status code %d", createJobResponseWithStatus.Status)
	}
}

func TestDeleteManyJobsSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var deleteManyJobsResponse DeleteManyJobsResponse
	if err := gofakeit.Struct(&deleteManyJobsResponse); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&deleteManyJobsResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Delete(jobsPath).
		Reply(200).
		JSON(&deleteManyJobsResponse)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	deleteManyJobsResponseWithStatus, err := client.DeleteManyJobs(ctx, &DeleteManyJobsRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if deleteManyJobsResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", deleteManyJobsResponseWithStatus.Status)
	}

	if diff := deep.Equal(deleteManyJobsResponse, deleteManyJobsResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestDeleteManyJobsFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	defer gock.Off()

	var deleteManyJobsResponse DeleteManyJobsResponse
	if err := gofakeit.Struct(&deleteManyJobsResponse); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&deleteManyJobsResponse)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Delete(jobsPath).
		Reply(400).
		JSON(&deleteManyJobsResponse)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	deleteManyJobsResponseWithStatus, err := client.DeleteManyJobs(ctx, &DeleteManyJobsRequest{})
	if err != nil {
		t.Fatal(err)
	}

	if deleteManyJobsResponseWithStatus.Status != 400 {
		t.Fatalf("unexpected response code: %d", deleteManyJobsResponseWithStatus.Status)
	}
}

func newGock() *gock.Request {
	return gock.New(myFakeEndpoint).
		MatchHeader("Authorization", fmt.Sprintf("apiKey %s", myFakeAPIKey)).
		MatchHeader("Content-Type", "application/json")
}

func TestGetJobSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	var jobResponseMock GetJobResponse
	if err := gofakeit.Struct(&jobResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(fmt.Sprintf("%s/some-id", jobsPath)).
		Reply(200).
		JSON(&jobResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobResponseWithStatus, err := client.GetJob(ctx, &GetJobRequest{
		ID: "some-id",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobResponseMock, jobResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestGetJobFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	var jobResponseMock GetJobResponse
	if err := gofakeit.Struct(&jobResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(fmt.Sprintf("%s/some-id", jobsPath)).
		Reply(400).
		JSON(&jobResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobResponseWithStatus, err := client.GetJob(ctx, &GetJobRequest{
		ID: "some-id",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobResponseWithStatus.Status != 400 {
		t.Fatalf("unexpected status %d", jobResponseWithStatus.Status)
	}
}

func TestDeleteJobSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	var jobResponseMock DeleteJobResponse
	if err := gofakeit.Struct(&jobResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Delete(fmt.Sprintf("%s/some-id", jobsPath)).
		Reply(200).
		JSON(&jobResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobResponseWithStatus, err := client.DeleteJob(ctx, &DeleteJobRequest{
		ID: "some-id",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobResponseMock, jobResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestDeleteJobFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	var jobResponseMock DeleteJobResponse
	if err := gofakeit.Struct(&jobResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Delete(fmt.Sprintf("%s/some-id", jobsPath)).
		Reply(400).
		JSON(&jobResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobResponseWithStatus, err := client.DeleteJob(ctx, &DeleteJobRequest{
		ID: "some-id",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobResponseWithStatus.Status != 400 {
		t.Fatalf("unexpected status %d", jobResponseWithStatus.Status)
	}
}

func TestGetJobOutputSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	jobResponseMock := GetJobOutputResponse{
		"0": 0.2,
		"1": 0.4,
		"2": 0.4,
	}

	mockJson, err := json.Marshal(&jobResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Get(fmt.Sprintf("%s/some-id/results", jobsPath)).
		Reply(200).
		JSON(&jobResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobResponseWithStatus, err := client.GetJobOutput(ctx, &GetJobOutputRequest{
		ID: "some-id",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobResponseMock, jobResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}

func TestCancelJobSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	defer gock.Off()

	jobResponseMock := CancelJobResponse{
		ID:     "some-id",
		Status: "canceled",
	}

	mockJson, err := json.Marshal(&jobResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	newGock().
		Put(fmt.Sprintf("%s/some-id/status/cancel", jobsPath)).
		Reply(200).
		JSON(&jobResponseMock)

	client := NewClient(myFakeEndpoint, myFakeAPIKey)
	jobResponseWithStatus, err := client.CancelJob(ctx, &CancelJobRequest{
		ID: "some-id",
	})
	if err != nil {
		t.Fatal(err)
	}

	if jobResponseWithStatus.Status != http.StatusOK {
		t.Fatalf("unexpected status: %d", jobResponseWithStatus.Status)
	}

	if diff := deep.Equal(jobResponseMock, jobResponseWithStatus.Response); len(diff) > 0 {
		t.Fatalf("unexpected diff: %s", diff)
	}
}
