package cronjoborg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetJobsResponse struct {
	Jobs []Job `json:"jobs,omitempty"`
}

func (c *Client) GetJobs() ([]Job, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/jobs", c.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	response := GetJobsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Jobs, nil
}
