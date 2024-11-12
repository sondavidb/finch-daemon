// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch-daemon/api/handlers/system (interfaces: Service)

// Package mocks_system is a generated GoMock package.
package mocks_system

import (
	context "context"
	reflect "reflect"

	config "github.com/containerd/nerdctl/v2/pkg/config"
	dockercompat "github.com/containerd/nerdctl/v2/pkg/inspecttypes/dockercompat"
	gomock "github.com/golang/mock/gomock"
	events "github.com/runfinch/finch-daemon/api/events"
	types "github.com/runfinch/finch-daemon/api/types"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockService) Auth(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockServiceMockRecorder) Auth(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockService)(nil).Auth), arg0, arg1, arg2, arg3)
}

// GetInfo mocks base method.
func (m *MockService) GetInfo(arg0 context.Context, arg1 *config.Config) (*dockercompat.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", arg0, arg1)
	ret0, _ := ret[0].(*dockercompat.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockServiceMockRecorder) GetInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockService)(nil).GetInfo), arg0, arg1)
}

// GetVersion mocks base method.
func (m *MockService) GetVersion(arg0 context.Context) (*types.VersionInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0)
	ret0, _ := ret[0].(*types.VersionInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockServiceMockRecorder) GetVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockService)(nil).GetVersion), arg0)
}

// SubscribeEvents mocks base method.
func (m *MockService) SubscribeEvents(arg0 context.Context, arg1 map[string][]string) (<-chan *events.Event, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeEvents", arg0, arg1)
	ret0, _ := ret[0].(<-chan *events.Event)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// SubscribeEvents indicates an expected call of SubscribeEvents.
func (mr *MockServiceMockRecorder) SubscribeEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEvents", reflect.TypeOf((*MockService)(nil).SubscribeEvents), arg0, arg1)
}
