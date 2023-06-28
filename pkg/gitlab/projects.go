package gitlab

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

type GitLabProject struct {
	Id              int
	Name            string
	RegistryEnabled bool   `json:"container_registry_enabled"`
	RegistryPrefix  string `json:"container_registry_image_prefix"`
}

func (app *Application) GetProject() (*GitLabProject, error) {
	var err error
	var resp *resty.Response
	project := &GitLabProject{}
	resp, err = app.client.R().
		SetPathParams(map[string]string{
			"projectId": strconv.Itoa(app.config.projectId),
		}).
		SetResult(project).
		Get("projects/{projectId}")
	if err != nil {
		err = fmt.Errorf("failed to retrieve GitLab project: %s", err)
	} else if resp.StatusCode() == 401 {
		err = errors.New("not authorized to list GitLab project")
	}
	return project, err
}