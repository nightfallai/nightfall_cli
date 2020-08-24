// Code generated by MockGen. DO NOT EDIT.
// Source: ../githubintf/github_client.go

// Package githubclient_mock is a generated GoMock package.
package githubclient_mock

import (
	gomock "github.com/golang/mock/gomock"
	githubintf "github.com/nightfallai/nightfall_code_scanner/internal/interfaces/githubintf"
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
