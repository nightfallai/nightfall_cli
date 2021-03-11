// Code generated by MockGen. DO NOT EDIT.
// Source: ../githubintf/github_checks.go

// Package githubchecks_mock is a generated GoMock package.
package githubchecks_mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v33/github"
	reflect "reflect"
)

// GithubChecks is a mock of GithubChecks interface
type GithubChecks struct {
	ctrl     *gomock.Controller
	recorder *GithubChecksMockRecorder
}

// GithubChecksMockRecorder is the mock recorder for GithubChecks
type GithubChecksMockRecorder struct {
	mock *GithubChecks
}

// NewGithubChecks creates a new mock instance
func NewGithubChecks(ctrl *gomock.Controller) *GithubChecks {
	mock := &GithubChecks{ctrl: ctrl}
	mock.recorder = &GithubChecksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *GithubChecks) EXPECT() *GithubChecksMockRecorder {
	return m.recorder
}

// CreateCheckRun mocks base method
func (m *GithubChecks) CreateCheckRun(ctx context.Context, owner, repo string, opts github.CreateCheckRunOptions) (*github.CheckRun, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCheckRun", ctx, owner, repo, opts)
	ret0, _ := ret[0].(*github.CheckRun)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCheckRun indicates an expected call of CreateCheckRun
func (mr *GithubChecksMockRecorder) CreateCheckRun(ctx, owner, repo, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCheckRun", reflect.TypeOf((*GithubChecks)(nil).CreateCheckRun), ctx, owner, repo, opts)
}

// UpdateCheckRun mocks base method
func (m *GithubChecks) UpdateCheckRun(ctx context.Context, owner, repo string, checkRunID int64, opts github.UpdateCheckRunOptions) (*github.CheckRun, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCheckRun", ctx, owner, repo, checkRunID, opts)
	ret0, _ := ret[0].(*github.CheckRun)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateCheckRun indicates an expected call of UpdateCheckRun
func (mr *GithubChecksMockRecorder) UpdateCheckRun(ctx, owner, repo, checkRunID, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCheckRun", reflect.TypeOf((*GithubChecks)(nil).UpdateCheckRun), ctx, owner, repo, checkRunID, opts)
}
