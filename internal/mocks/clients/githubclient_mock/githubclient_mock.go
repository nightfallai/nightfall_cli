// Code generated by MockGen. DO NOT EDIT.
// Source: ../githubintf/github_client.go

// Package githubclient_mock is a generated GoMock package.
package githubclient_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v31/github"
	githubintf "github.com/watchtowerai/nightfall_dlp/internal/interfaces/githubintf"
	http "net/http"
	reflect "reflect"
)

// GithubClient is a mock of GithubClient interface
type GithubClient struct {
	ctrl     *gomock.Controller
	recorder *GithubClientMockRecorder
}

// GithubClientMockRecorder is the mock recorder for GithubClient
type GithubClientMockRecorder struct {
	mock *GithubClient
}

// NewGithubClient creates a new mock instance
func NewGithubClient(ctrl *gomock.Controller) *GithubClient {
	mock := &GithubClient{ctrl: ctrl}
	mock.recorder = &GithubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *GithubClient) EXPECT() *GithubClientMockRecorder {
	return m.recorder
}

// PullRequestService mocks base method
func (m *GithubClient) PullRequestService() githubintf.GithubPullRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullRequestService")
	ret0, _ := ret[0].(githubintf.GithubPullRequest)
	return ret0
}

// PullRequestService indicates an expected call of PullRequestService
func (mr *GithubClientMockRecorder) PullRequestService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullRequestService", reflect.TypeOf((*GithubClient)(nil).PullRequestService))
}

// ChecksService mocks base method
func (m *GithubClient) ChecksService() githubintf.GithubChecks {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChecksService")
	ret0, _ := ret[0].(githubintf.GithubChecks)
	return ret0
}

// ChecksService indicates an expected call of ChecksService
func (mr *GithubClientMockRecorder) ChecksService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChecksService", reflect.TypeOf((*GithubClient)(nil).ChecksService))
}

// Do mocks base method
func (m *GithubClient) Do(ctx context.Context, req *http.Request, v interface{}) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", ctx, req, v)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *GithubClientMockRecorder) Do(ctx, req, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*GithubClient)(nil).Do), ctx, req, v)
}

// GetRawBySha mocks base method
func (m *GithubClient) GetRawBySha(ctx context.Context, owner, repo, base, headOrSha string) (string, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawBySha", ctx, owner, repo, base, headOrSha)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRawBySha indicates an expected call of GetRawBySha
func (mr *GithubClientMockRecorder) GetRawBySha(ctx, owner, repo, base, headOrSha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawBySha", reflect.TypeOf((*GithubClient)(nil).GetRawBySha), ctx, owner, repo, base, headOrSha)
}