// Code generated by MockGen. DO NOT EDIT.
// Source: ../logger/logger.go

// Package logger_mock is a generated GoMock package.
package logger_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Logger is a mock of Logger interface
type Logger struct {
	ctrl     *gomock.Controller
	recorder *LoggerMockRecorder
}

// LoggerMockRecorder is the mock recorder for Logger
type LoggerMockRecorder struct {
	mock *Logger
}

// NewLogger creates a new mock instance
func NewLogger(ctrl *gomock.Controller) *Logger {
	mock := &Logger{ctrl: ctrl}
	mock.recorder = &LoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Logger) EXPECT() *LoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *Logger) Debug(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Debug", msg)
}

// Debug indicates an expected call of Debug
func (mr *LoggerMockRecorder) Debug(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*Logger)(nil).Debug), msg)
}

// Info mocks base method
func (m *Logger) Info(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", msg)
}

// Info indicates an expected call of Info
func (mr *LoggerMockRecorder) Info(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*Logger)(nil).Info), msg)
}

// Warning mocks base method
func (m *Logger) Warning(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Warning", msg)
}

// Warning indicates an expected call of Warning
func (mr *LoggerMockRecorder) Warning(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*Logger)(nil).Warning), msg)
}

// Error mocks base method
func (m *Logger) Error(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", msg)
}

// Error indicates an expected call of Error
func (mr *LoggerMockRecorder) Error(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*Logger)(nil).Error), msg)
}