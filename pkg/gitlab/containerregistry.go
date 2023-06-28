package gitlab

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type RegistryRepository struct{
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	ProjectId int    `json:"project_id"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
}

type RegistryRepositories []RegistryRepository

func (app *Application) GetRegistryRepositories() (*RegistryRepositories, error) {
	var err error
	var resp *resty.Response
	repos := &RegistryRepositories{}
	resp, err = app.client.R().
		SetPathParams(map[string]string{
			"projectId": strconv.Itoa(app.config.projectId),
		}).
		SetResult(repos).
		Get("projects/{projectId}/registry/repositories")
	if err != nil {
		err = fmt.Errorf("failed to retrieve GitLab project repositories: %s", err)
	} else if resp.StatusCode() == 401 {
		err = errors.New("not authorized to list GitLab project")
	}
	return repos, err
}

func (repos *RegistryRepositories) PrintRegistryRepositories() {
	for _, j := range *repos {
		fmt.Printf("%+v\n", j)
	}
}