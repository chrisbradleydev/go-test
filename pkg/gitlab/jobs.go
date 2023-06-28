package gitlab

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type Job struct{
	Id     int    `json:"id"`
	Status string `json:"status"`
	Stage  string `json:"stage"`
	Name   string `json:"name"`
    Ref    string `json:"ref"`
    Tag    bool   `json:"tag"`
}

type Jobs []Job

func (app *Application) GetJobs() (*Jobs, error) {
	var err error
	var resp *resty.Response
	jobs := &Jobs{}
	resp, err = app.client.R().
		SetPathParams(map[string]string{
			"projectId": strconv.Itoa(app.config.projectId),
		}).
		SetResult(jobs).
		SetQueryParam("scope[]", "success").
		Get("/projects/{projectId}/jobs")
	if err != nil {
		err = fmt.Errorf("failed to retrieve GitLab project jobs: %s", err)
	} else if resp.StatusCode() == 401 {
		err = errors.New("not authorized to list GitLab project")
	}
	return jobs, err
}

func (jobs *Jobs) PrintJobs() {
	for _, j := range *jobs {
		fmt.Printf("%+v\n", j)
	}
}