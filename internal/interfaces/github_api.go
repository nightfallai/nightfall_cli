package interfaces

import (
	"context"

	"github.com/google/go-github/v31/github"
)

//go:generate go run github.com/golang/mock/mockgen -destination=../mocks/clients/githubapi_mock/api_mock.go -source=../interfaces/github_api.go -package=githubapi_mock -mock_names=GithubAPI=GithubAPI

type GithubAPI interface {
	GetRaw(ctx context.Context, owner string, repo string, number int, opts github.RawOptions) (string, *github.Response, error)
	GetRawBySha(ctx context.Context, owner string, repo string, base string, headOrSha string) (string, *github.Response, error)
	ListCheckRunsForRef(ctx context.Context, owner, repo, ref string, opts *github.ListCheckRunsOptions) (*github.ListCheckRunsResults, *github.Response, error)
	CreateCheckRun(ctx context.Context, owner, repo string, opts github.CreateCheckRunOptions) (*github.CheckRun, *github.Response, error)
	UpdateCheckRun(ctx context.Context, owner, repo string, checkRunID int64, opts github.UpdateCheckRunOptions) (*github.CheckRun, *github.Response, error)
}