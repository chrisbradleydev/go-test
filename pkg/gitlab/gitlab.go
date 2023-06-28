package gitlab

import (
	"flag"
	"time"

	"github.com/go-resty/resty/v2"
)

type config struct {
	apiPath   string
	apiToken  string
	projectId int
}

type Application struct {
	client    *resty.Client
	config    config
}

func NewApp() *Application {
	var cfg config
	flag.StringVar(&cfg.apiPath, "gitlab-api-path", "", "GitLab API path")
	flag.StringVar(&cfg.apiToken, "gitlab-api-token", "", "GitLab API token")
	flag.IntVar(&cfg.projectId, "gitlab-project-id", 1, "GitLab project id")
	flag.Parse()

	client := resty.New().SetBaseURL(cfg.apiPath).SetTimeout(10*time.Second)
	client.SetHeader("PRIVATE-TOKEN", cfg.apiToken)
	return &Application{
		client: client,
		config: cfg,
	}
}