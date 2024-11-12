// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch-daemon/internal/backend (interfaces: NerdctlImageSvc)

// Package mocks_backend is a generated GoMock package.
package mocks_backend

import (
	context "context"
	io "io"
	reflect "reflect"

	images "github.com/containerd/containerd/images"
	platforms "github.com/containerd/platforms"
	remotes "github.com/containerd/containerd/remotes"
	docker "github.com/containerd/containerd/remotes/docker"
	imgutil "github.com/containerd/nerdctl/v2/pkg/imgutil"
	dockerconfigresolver "github.com/containerd/nerdctl/v2/pkg/imgutil/dockerconfigresolver"
	dockercompat "github.com/containerd/nerdctl/v2/pkg/inspecttypes/dockercompat"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// MockNerdctlImageSvc is a mock of NerdctlImageSvc interface.
type MockNerdctlImageSvc struct {
	ctrl     *gomock.Controller
	recorder *MockNerdctlImageSvcMockRecorder
}

// MockNerdctlImageSvcMockRecorder is the mock recorder for MockNerdctlImageSvc.
type MockNerdctlImageSvcMockRecorder struct {
	mock *MockNerdctlImageSvc
}

// NewMockNerdctlImageSvc creates a new mock instance.
func NewMockNerdctlImageSvc(ctrl *gomock.Controller) *MockNerdctlImageSvc {
	mock := &MockNerdctlImageSvc{ctrl: ctrl}
	mock.recorder = &MockNerdctlImageSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNerdctlImageSvc) EXPECT() *MockNerdctlImageSvcMockRecorder {
	return m.recorder
}

// GetDataStore mocks base method.
func (m *MockNerdctlImageSvc) GetDataStore() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataStore")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataStore indicates an expected call of GetDataStore.
func (mr *MockNerdctlImageSvcMockRecorder) GetDataStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataStore", reflect.TypeOf((*MockNerdctlImageSvc)(nil).GetDataStore))
}

// GetDockerResolver mocks base method.
func (m *MockNerdctlImageSvc) GetDockerResolver(arg0 context.Context, arg1 string, arg2 dockerconfigresolver.AuthCreds) (remotes.Resolver, docker.StatusTracker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDockerResolver", arg0, arg1, arg2)
	ret0, _ := ret[0].(remotes.Resolver)
	ret1, _ := ret[1].(docker.StatusTracker)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDockerResolver indicates an expected call of GetDockerResolver.
func (mr *MockNerdctlImageSvcMockRecorder) GetDockerResolver(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDockerResolver", reflect.TypeOf((*MockNerdctlImageSvc)(nil).GetDockerResolver), arg0, arg1, arg2)
}

// InspectImage mocks base method.
func (m *MockNerdctlImageSvc) InspectImage(arg0 context.Context, arg1 images.Image) (*dockercompat.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectImage", arg0, arg1)
	ret0, _ := ret[0].(*dockercompat.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImage indicates an expected call of InspectImage.
func (mr *MockNerdctlImageSvcMockRecorder) InspectImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectImage", reflect.TypeOf((*MockNerdctlImageSvc)(nil).InspectImage), arg0, arg1)
}

// LoadImage mocks base method.
func (m *MockNerdctlImageSvc) LoadImage(arg0 context.Context, arg1 string, arg2 io.Writer, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadImage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadImage indicates an expected call of LoadImage.
func (mr *MockNerdctlImageSvcMockRecorder) LoadImage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*MockNerdctlImageSvc)(nil).LoadImage), arg0, arg1, arg2, arg3)
}

// PullImage mocks base method.
func (m *MockNerdctlImageSvc) PullImage(arg0 context.Context, arg1, arg2 io.Writer, arg3 remotes.Resolver, arg4 string, arg5 []v1.Platform) (*imgutil.EnsuredImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*imgutil.EnsuredImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullImage indicates an expected call of PullImage.
func (mr *MockNerdctlImageSvcMockRecorder) PullImage(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockNerdctlImageSvc)(nil).PullImage), arg0, arg1, arg2, arg3, arg4, arg5)
}

// PushImage mocks base method.
func (m *MockNerdctlImageSvc) PushImage(arg0 context.Context, arg1 remotes.Resolver, arg2 docker.StatusTracker, arg3 io.Writer, arg4, arg5 string, arg6 platforms.MatchComparer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushImage", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushImage indicates an expected call of PushImage.
func (mr *MockNerdctlImageSvcMockRecorder) PushImage(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushImage", reflect.TypeOf((*MockNerdctlImageSvc)(nil).PushImage), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SearchImage mocks base method.
func (m *MockNerdctlImageSvc) SearchImage(arg0 context.Context, arg1 string) (int, int, []*images.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchImage", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].([]*images.Image)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SearchImage indicates an expected call of SearchImage.
func (mr *MockNerdctlImageSvcMockRecorder) SearchImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchImage", reflect.TypeOf((*MockNerdctlImageSvc)(nil).SearchImage), arg0, arg1)
}
