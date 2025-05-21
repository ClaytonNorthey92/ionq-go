package ionq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

const jobsPath = "jobs"

type Job struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Target string `json:"target,omitempty"`
	Noise  struct {
		Model string `json:"model,omitempty"`
		Seed  int    `json:"seed,omitempty"`
	} `json:"noise,omitempty"`
	Metadata struct {
		CustomKey string `json:"custom_key,omitempty"`
	} `json:"metadata,omitempty"`
	Shots           int `json:"shots,omitempty"`
	ErrorMitigation struct {
		Debias bool `json:"debias,omitempty"`
	} `json:"error_mitigation,omitempty"`
	GateCounts struct {
		OneQ int `json:"1q,omitempty"`
		TwoQ int `json:"2q,omitempty"`
	} `json:"gate_counts,omitempty"`
	Qubits                 int      `json:"qubits,omitempty"`
	CostUsd                float64  `json:"cost_usd,omitempty"`
	Request                int      `json:"request,omitempty"`
	Start                  int      `json:"start,omitempty"`
	Response               int      `json:"response,omitempty"`
	ExecutionTime          int      `json:"execution_time,omitempty"`
	PredictedExecutionTime int      `json:"predicted_execution_time,omitempty"`
	Children               []string `json:"children,omitempty"`
	ResultsURL             string   `json:"results_url,omitempty"`
	Failure                struct {
		Error string `json:"error,omitempty"`
		Code  string `json:"code,omitempty"`
	} `json:"failure,omitempty"`
	Warning struct {
		Messages []string `json:"messages,omitempty"`
	} `json:"warning,omitempty"`
	Circuits int `json:"circuits,omitempty"`
}

type GetJobsResponse struct {
	Jobs []Job  `json:"jobs,omitempty"`
	Next string `json:"next,omitempty"`
}

type GetJobsRequest struct {
	IDs    []string `url:"id"`
	Status string   `url:"status"`
	Limit  uint     `url:"limit"`
	Next   string   `url:"next"`
}

type GetJobRequest struct {
	ID string `url:"id"`
}

type GetJobOutputRequest GetJobRequest

// these have the same structure
type DeleteJobRequest GetJobRequest
type CancelJobRequest GetJobRequest

type DeleteJobResponseWithStatus struct {
	Response DeleteJobResponse
	Status   int
}

type GetJobResponse Job

type GetJobResponseWithStatus struct {
	Response GetJobResponse
	Status   int
}

type GetJobsResponseWithStatus struct {
	Response GetJobsResponse
	Status   int
}

type CircuitInput struct {
	Gate     string `json:"gate,omitempty"`
	Target   *uint  `json:"target,omitempty"`
	Targets  []uint `json:"targets,omitempty`
	Control  uint   `json:"control,omitempty"`
	Controls []uint `json:"controls,omitempty"`
	Rotation int    `json:"rotation,omitempty"`
}

type JobInput struct {
	Circuit []CircuitInput `json:"circuit,omitempty"`
	Qubits  uint           `json:"qubits"`
	Format  string         `json:"format,omitempty"`
	Gateset string         `json:"gateset,omitempty"`
}

type NoiseInput struct {
	Model string `json:"model,omitempty"`
	Seed  int    `json:"seed,omitempty"`
}

type ErrorMitigationInput struct {
	Debias bool `json:"debias,omitempty"`
}

type CreateJobRequest struct {
	Name            string                `json:"name,omitempty"`
	Metadata        map[string]string     `json:"metadata,omitempty"`
	Shots           uint                  `json:"shots,omitempty"`
	Target          string                `json:"target,omitempty"`
	Noise           *NoiseInput           `json:"noise,omitempty"`
	Input           JobInput              `json:"input,omitempty"`
	ErrorMitigation *ErrorMitigationInput `json:"error_mitigation,omitempty"`
}

type CreateJobResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CreateJobResponseWithStatus struct {
	Response CreateJobResponse
	Status   int
}

type DeleteManyJobsRequest struct {
	IDs []string `json:"ids"`
}

type DeleteManyJobsResponse struct {
	IDS    []string `json:"ids"`
	Status string   `json:"status"`
}

type DeleteJobResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CancelJobResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type CancelJobResponseWithStatus struct {
	Response CancelJobResponse
	Status   int
}

type DeleteManyJobsResponseWithStatus struct {
	Response DeleteManyJobsResponse
	Status   int
}

type GetJobOutputResponse map[string]float32

type GetJobOutputResponseWithStatus struct {
	Response GetJobOutputResponse
	Status   int
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("apiKey %s", c.apiKey))
}

func (c *Client) GetJobs(ctx context.Context, getJobsRequest *GetJobsRequest) (*GetJobsResponseWithStatus, error) {
	url := c.makeURL(jobsPath)

	v, err := query.Values(&getJobsRequest)
	if err != nil {
		return nil, err
	}

	url += fmt.Sprintf("?%s", v.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jobsResponse GetJobsResponse
	if err := json.Unmarshal(body, &jobsResponse); err != nil {
		return nil, err
	}

	return &GetJobsResponseWithStatus{
		Response: jobsResponse,
		Status:   res.StatusCode,
	}, nil
}

func (c *Client) CreateJob(ctx context.Context, createJobRequest *CreateJobRequest) (*CreateJobResponseWithStatus, error) {
	url := c.makeURL(jobsPath)

	reqBody, err := json.Marshal(createJobRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var createJobResponseWithStatus CreateJobResponseWithStatus

	if err := json.Unmarshal(body, &createJobResponseWithStatus.Response); err != nil {
		return nil, err
	}

	createJobResponseWithStatus.Status = res.StatusCode

	return &createJobResponseWithStatus, nil
}

func (c *Client) DeleteManyJobs(ctx context.Context, deleteManyJobsRequest *DeleteManyJobsRequest) (*DeleteManyJobsResponseWithStatus, error) {
	url := c.makeURL(jobsPath)

	reqBody, err := json.Marshal(deleteManyJobsRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var deleteManyJobsResponseWithStatus DeleteManyJobsResponseWithStatus

	if err := json.Unmarshal(body, &deleteManyJobsResponseWithStatus.Response); err != nil {
		return nil, err
	}

	deleteManyJobsResponseWithStatus.Status = res.StatusCode

	return &deleteManyJobsResponseWithStatus, nil
}

func (c *Client) GetJob(ctx context.Context, getJobRequest *GetJobRequest) (*GetJobResponseWithStatus, error) {
	url := c.makeURL(fmt.Sprintf("%s/%s", jobsPath, getJobRequest.ID))

	// TODO: add "include" and "exclude" query params

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var getJobResponse GetJobResponse
	if err := json.Unmarshal(body, &getJobResponse); err != nil {
		return nil, err
	}

	return &GetJobResponseWithStatus{
		Response: getJobResponse,
		Status:   res.StatusCode,
	}, nil
}

func (c *Client) GetJobOutput(ctx context.Context, getJobOutputRequest *GetJobOutputRequest) (*GetJobOutputResponseWithStatus, error) {
	url := c.makeURL(fmt.Sprintf("%s/%s/results", jobsPath, getJobOutputRequest.ID))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var getJobOutputResponse GetJobOutputResponse
	if err := json.Unmarshal(body, &getJobOutputResponse); err != nil {
		return nil, err
	}

	return &GetJobOutputResponseWithStatus{
		Response: getJobOutputResponse,
		Status:   res.StatusCode,
	}, nil
}

func (c *Client) DeleteJob(ctx context.Context, deleteJobRequest *DeleteJobRequest) (*DeleteJobResponseWithStatus, error) {
	url := c.makeURL(fmt.Sprintf("%s/%s", jobsPath, deleteJobRequest.ID))

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var deleteJobResponse DeleteJobResponse
	if err := json.Unmarshal(body, &deleteJobResponse); err != nil {
		return nil, err
	}

	return &DeleteJobResponseWithStatus{
		Response: deleteJobResponse,
		Status:   res.StatusCode,
	}, nil
}

func (c *Client) CancelJob(ctx context.Context, cancelJobRequest *CancelJobRequest) (*CancelJobResponseWithStatus, error) {
	url := c.makeURL(fmt.Sprintf("%s/%s/status/cancel", jobsPath, cancelJobRequest.ID))

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cancelJobResponse CancelJobResponse
	if err := json.Unmarshal(body, &cancelJobResponse); err != nil {
		return nil, err
	}

	return &CancelJobResponseWithStatus{
		Response: cancelJobResponse,
		Status:   res.StatusCode,
	}, nil
}
