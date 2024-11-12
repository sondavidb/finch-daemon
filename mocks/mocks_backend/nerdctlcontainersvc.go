// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/runfinch/finch-daemon/internal/backend (interfaces: NerdctlContainerSvc)

// Package mocks_backend is a generated GoMock package.
package mocks_backend

import (
	context "context"
	io "io"
	os "os"
	reflect "reflect"
	time "time"

	containerd "github.com/containerd/containerd"
	types "github.com/containerd/nerdctl/v2/pkg/api/types"
	container "github.com/containerd/nerdctl/v2/pkg/cmd/container"
	containerutil "github.com/containerd/nerdctl/v2/pkg/containerutil"
	dockercompat "github.com/containerd/nerdctl/v2/pkg/inspecttypes/dockercompat"
	native "github.com/containerd/nerdctl/v2/pkg/inspecttypes/native"
	logging "github.com/containerd/nerdctl/v2/pkg/logging"
	gomock "github.com/golang/mock/gomock"
)

// MockNerdctlContainerSvc is a mock of NerdctlContainerSvc interface.
type MockNerdctlContainerSvc struct {
	ctrl     *gomock.Controller
	recorder *MockNerdctlContainerSvcMockRecorder
}

// MockNerdctlContainerSvcMockRecorder is the mock recorder for MockNerdctlContainerSvc.
type MockNerdctlContainerSvcMockRecorder struct {
	mock *MockNerdctlContainerSvc
}

// NewMockNerdctlContainerSvc creates a new mock instance.
func NewMockNerdctlContainerSvc(ctrl *gomock.Controller) *MockNerdctlContainerSvc {
	mock := &MockNerdctlContainerSvc{ctrl: ctrl}
	mock.recorder = &MockNerdctlContainerSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNerdctlContainerSvc) EXPECT() *MockNerdctlContainerSvcMockRecorder {
	return m.recorder
}

// CreateContainer mocks base method.
func (m *MockNerdctlContainerSvc) CreateContainer(arg0 context.Context, arg1 []string, arg2 containerutil.NetworkOptionsManager, arg3 types.ContainerCreateOptions) (containerd.Container, func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(containerd.Container)
	ret1, _ := ret[1].(func())
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockNerdctlContainerSvcMockRecorder) CreateContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).CreateContainer), arg0, arg1, arg2, arg3)
}

// GetDataStore mocks base method.
func (m *MockNerdctlContainerSvc) GetDataStore() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataStore")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataStore indicates an expected call of GetDataStore.
func (mr *MockNerdctlContainerSvcMockRecorder) GetDataStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataStore", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).GetDataStore))
}

// GetNerdctlExe mocks base method.
func (m *MockNerdctlContainerSvc) GetNerdctlExe() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNerdctlExe")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNerdctlExe indicates an expected call of GetNerdctlExe.
func (mr *MockNerdctlContainerSvcMockRecorder) GetNerdctlExe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNerdctlExe", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).GetNerdctlExe))
}

// InspectContainer mocks base method.
func (m *MockNerdctlContainerSvc) InspectContainer(arg0 context.Context, arg1 containerd.Container) (*dockercompat.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectContainer", arg0, arg1)
	ret0, _ := ret[0].(*dockercompat.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer.
func (mr *MockNerdctlContainerSvcMockRecorder) InspectContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectContainer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).InspectContainer), arg0, arg1)
}

// InspectNetNS mocks base method.
func (m *MockNerdctlContainerSvc) InspectNetNS(arg0 context.Context, arg1 int) (*native.NetNS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectNetNS", arg0, arg1)
	ret0, _ := ret[0].(*native.NetNS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectNetNS indicates an expected call of InspectNetNS.
func (mr *MockNerdctlContainerSvcMockRecorder) InspectNetNS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectNetNS", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).InspectNetNS), arg0, arg1)
}

// ListContainers mocks base method.
func (m *MockNerdctlContainerSvc) ListContainers(arg0 context.Context, arg1 types.ContainerListOptions) ([]container.ListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers", arg0, arg1)
	ret0, _ := ret[0].([]container.ListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockNerdctlContainerSvcMockRecorder) ListContainers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).ListContainers), arg0, arg1)
}

// LoggingInitContainerLogViewer mocks base method.
func (m *MockNerdctlContainerSvc) LoggingInitContainerLogViewer(arg0 map[string]string, arg1 logging.LogViewOptions, arg2 chan os.Signal, arg3 bool) (*logging.ContainerLogViewer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoggingInitContainerLogViewer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*logging.ContainerLogViewer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoggingInitContainerLogViewer indicates an expected call of LoggingInitContainerLogViewer.
func (mr *MockNerdctlContainerSvcMockRecorder) LoggingInitContainerLogViewer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoggingInitContainerLogViewer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).LoggingInitContainerLogViewer), arg0, arg1, arg2, arg3)
}

// LoggingPrintLogsTo mocks base method.
func (m *MockNerdctlContainerSvc) LoggingPrintLogsTo(arg0, arg1 io.Writer, arg2 *logging.ContainerLogViewer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoggingPrintLogsTo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoggingPrintLogsTo indicates an expected call of LoggingPrintLogsTo.
func (mr *MockNerdctlContainerSvcMockRecorder) LoggingPrintLogsTo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoggingPrintLogsTo", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).LoggingPrintLogsTo), arg0, arg1, arg2)
}

// NewNetworkingOptionsManager mocks base method.
func (m *MockNerdctlContainerSvc) NewNetworkingOptionsManager(arg0 types.NetworkOptions) (containerutil.NetworkOptionsManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNetworkingOptionsManager", arg0)
	ret0, _ := ret[0].(containerutil.NetworkOptionsManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNetworkingOptionsManager indicates an expected call of NewNetworkingOptionsManager.
func (mr *MockNerdctlContainerSvcMockRecorder) NewNetworkingOptionsManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNetworkingOptionsManager", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).NewNetworkingOptionsManager), arg0)
}

// RemoveContainer mocks base method.
func (m *MockNerdctlContainerSvc) RemoveContainer(arg0 context.Context, arg1 containerd.Container, arg2, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveContainer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer.
func (mr *MockNerdctlContainerSvcMockRecorder) RemoveContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).RemoveContainer), arg0, arg1, arg2, arg3)
}

// RenameContainer mocks base method.
func (m *MockNerdctlContainerSvc) RenameContainer(arg0 context.Context, arg1 containerd.Container, arg2 string, arg3 types.ContainerRenameOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameContainer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameContainer indicates an expected call of RenameContainer.
func (mr *MockNerdctlContainerSvcMockRecorder) RenameContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameContainer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).RenameContainer), arg0, arg1, arg2, arg3)
}

// StartContainer mocks base method.
func (m *MockNerdctlContainerSvc) StartContainer(arg0 context.Context, arg1 containerd.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainer indicates an expected call of StartContainer.
func (mr *MockNerdctlContainerSvcMockRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).StartContainer), arg0, arg1)
}

// StopContainer mocks base method.
func (m *MockNerdctlContainerSvc) StopContainer(arg0 context.Context, arg1 containerd.Container, arg2 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer.
func (mr *MockNerdctlContainerSvcMockRecorder) StopContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockNerdctlContainerSvc)(nil).StopContainer), arg0, arg1, arg2)
}
