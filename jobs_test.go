package ionq

import (
	"context"
	"encoding/json"
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

	var jobsResponseMock JobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	gock.New(myFakeEndpoint).
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

	var jobsResponseMock JobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	gock.New(myFakeEndpoint).
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

	var jobsResponseMock JobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	gock.New(myFakeEndpoint).
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

	var jobsResponseMock JobsResponse
	if err := gofakeit.Struct(&jobsResponseMock); err != nil {
		t.Fatal(err)
	}

	mockJson, err := json.Marshal(&jobsResponseMock)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("will mock response as: %s", mockJson)

	gock.New(myFakeEndpoint).
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
