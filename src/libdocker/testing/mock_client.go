// +build !dockerless

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Generated code, generated via: `mockgen -source src/libdocker/client.go  -destination src/libdocker/testing/mock_client.go -package testing`
// Edited by hand for boilerplate and gofmt.

// Package testing is a generated GoMock package.
package testing

import (
	libdocker "github.com/Mirantis/cri-dockerd/libdocker"
	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	image "github.com/docker/docker/api/types/image"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ListContainers mocks base method
func (m *MockInterface) ListContainers(
	options types.ContainerListOptions,
) ([]types.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers", options)
	ret0, _ := ret[0].([]types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers
func (mr *MockInterfaceMockRecorder) ListContainers(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ListContainers",
		reflect.TypeOf((*MockInterface)(nil).ListContainers),
		options,
	)
}

// InspectContainer mocks base method
func (m *MockInterface) InspectContainer(id string) (*types.ContainerJSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectContainer", id)
	ret0, _ := ret[0].(*types.ContainerJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer
func (mr *MockInterfaceMockRecorder) InspectContainer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"InspectContainer",
		reflect.TypeOf((*MockInterface)(nil).InspectContainer),
		id,
	)
}

// InspectContainerWithSize mocks base method
func (m *MockInterface) InspectContainerWithSize(id string) (*types.ContainerJSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectContainerWithSize", id)
	ret0, _ := ret[0].(*types.ContainerJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainerWithSize indicates an expected call of InspectContainerWithSize
func (mr *MockInterfaceMockRecorder) InspectContainerWithSize(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"InspectContainerWithSize",
		reflect.TypeOf((*MockInterface)(nil).InspectContainerWithSize),
		id,
	)
}

// CreateContainer mocks base method
func (m *MockInterface) CreateContainer(
	arg0 types.ContainerCreateConfig,
) (*container.ContainerCreateCreatedBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0)
	ret0, _ := ret[0].(*container.ContainerCreateCreatedBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockInterfaceMockRecorder) CreateContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"CreateContainer",
		reflect.TypeOf((*MockInterface)(nil).CreateContainer),
		arg0,
	)
}

// StartContainer mocks base method
func (m *MockInterface) StartContainer(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainer", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainer indicates an expected call of StartContainer
func (mr *MockInterfaceMockRecorder) StartContainer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"StartContainer",
		reflect.TypeOf((*MockInterface)(nil).StartContainer),
		id,
	)
}

// StopContainer mocks base method
func (m *MockInterface) StopContainer(id string, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", id, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (mr *MockInterfaceMockRecorder) StopContainer(id, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"StopContainer",
		reflect.TypeOf((*MockInterface)(nil).StopContainer),
		id,
		timeout,
	)
}

// UpdateContainerResources mocks base method
func (m *MockInterface) UpdateContainerResources(
	id string,
	updateConfig container.UpdateConfig,
) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContainerResources", id, updateConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContainerResources indicates an expected call of UpdateContainerResources
func (mr *MockInterfaceMockRecorder) UpdateContainerResources(
	id, updateConfig interface{},
) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"UpdateContainerResources",
		reflect.TypeOf((*MockInterface)(nil).UpdateContainerResources),
		id,
		updateConfig,
	)
}

// RemoveContainer mocks base method
func (m *MockInterface) RemoveContainer(id string, opts types.ContainerRemoveOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveContainer", id, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (mr *MockInterfaceMockRecorder) RemoveContainer(id, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"RemoveContainer",
		reflect.TypeOf((*MockInterface)(nil).RemoveContainer),
		id,
		opts,
	)
}

// InspectImageByRef mocks base method
func (m *MockInterface) InspectImageByRef(imageRef string) (*types.ImageInspect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectImageByRef", imageRef)
	ret0, _ := ret[0].(*types.ImageInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImageByRef indicates an expected call of InspectImageByRef
func (mr *MockInterfaceMockRecorder) InspectImageByRef(imageRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"InspectImageByRef",
		reflect.TypeOf((*MockInterface)(nil).InspectImageByRef),
		imageRef,
	)
}

// InspectImageByID mocks base method
func (m *MockInterface) InspectImageByID(imageID string) (*types.ImageInspect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectImageByID", imageID)
	ret0, _ := ret[0].(*types.ImageInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImageByID indicates an expected call of InspectImageByID
func (mr *MockInterfaceMockRecorder) InspectImageByID(imageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"InspectImageByID",
		reflect.TypeOf((*MockInterface)(nil).InspectImageByID),
		imageID,
	)
}

// ListImages mocks base method
func (m *MockInterface) ListImages(opts types.ImageListOptions) ([]types.ImageSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", opts)
	ret0, _ := ret[0].([]types.ImageSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages
func (mr *MockInterfaceMockRecorder) ListImages(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ListImages",
		reflect.TypeOf((*MockInterface)(nil).ListImages),
		opts,
	)
}

// PullImage mocks base method
func (m *MockInterface) PullImage(
	image string,
	auth types.AuthConfig,
	opts types.ImagePullOptions,
) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullImage", image, auth, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullImage indicates an expected call of PullImage
func (mr *MockInterfaceMockRecorder) PullImage(image, auth, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"PullImage",
		reflect.TypeOf((*MockInterface)(nil).PullImage),
		image,
		auth,
		opts,
	)
}

// RemoveImage mocks base method
func (m *MockInterface) RemoveImage(
	image string,
	opts types.ImageRemoveOptions,
) ([]types.ImageDeleteResponseItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImage", image, opts)
	ret0, _ := ret[0].([]types.ImageDeleteResponseItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveImage indicates an expected call of RemoveImage
func (mr *MockInterfaceMockRecorder) RemoveImage(image, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"RemoveImage",
		reflect.TypeOf((*MockInterface)(nil).RemoveImage),
		image,
		opts,
	)
}

// ImageHistory mocks base method
func (m *MockInterface) ImageHistory(id string) ([]image.HistoryResponseItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageHistory", id)
	ret0, _ := ret[0].([]image.HistoryResponseItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageHistory indicates an expected call of ImageHistory
func (mr *MockInterfaceMockRecorder) ImageHistory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ImageHistory",
		reflect.TypeOf((*MockInterface)(nil).ImageHistory),
		id,
	)
}

// Logs mocks base method
func (m *MockInterface) Logs(
	arg0 string,
	arg1 types.ContainerLogsOptions,
	arg2 libdocker.StreamOptions,
) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logs indicates an expected call of Logs
func (mr *MockInterfaceMockRecorder) Logs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Logs",
		reflect.TypeOf((*MockInterface)(nil).Logs),
		arg0,
		arg1,
		arg2,
	)
}

// Version mocks base method
func (m *MockInterface) Version() (*types.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*types.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockInterfaceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Version",
		reflect.TypeOf((*MockInterface)(nil).Version),
	)
}

// Info mocks base method
func (m *MockInterface) Info() (*types.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(*types.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockInterfaceMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Info",
		reflect.TypeOf((*MockInterface)(nil).Info),
	)
}

// CreateExec mocks base method
func (m *MockInterface) CreateExec(arg0 string, arg1 types.ExecConfig) (*types.IDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExec", arg0, arg1)
	ret0, _ := ret[0].(*types.IDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExec indicates an expected call of CreateExec
func (mr *MockInterfaceMockRecorder) CreateExec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"CreateExec",
		reflect.TypeOf((*MockInterface)(nil).CreateExec),
		arg0,
		arg1,
	)
}

// StartExec mocks base method
func (m *MockInterface) StartExec(
	arg0 string,
	arg1 types.ExecStartCheck,
	arg2 libdocker.StreamOptions,
) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartExec", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartExec indicates an expected call of StartExec
func (mr *MockInterfaceMockRecorder) StartExec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"StartExec",
		reflect.TypeOf((*MockInterface)(nil).StartExec),
		arg0,
		arg1,
		arg2,
	)
}

// InspectExec mocks base method
func (m *MockInterface) InspectExec(id string) (*types.ContainerExecInspect, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectExec", id)
	ret0, _ := ret[0].(*types.ContainerExecInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectExec indicates an expected call of InspectExec
func (mr *MockInterfaceMockRecorder) InspectExec(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"InspectExec",
		reflect.TypeOf((*MockInterface)(nil).InspectExec),
		id,
	)
}

// AttachToContainer mocks base method
func (m *MockInterface) AttachToContainer(
	arg0 string,
	arg1 types.ContainerAttachOptions,
	arg2 libdocker.StreamOptions,
) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachToContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachToContainer indicates an expected call of AttachToContainer
func (mr *MockInterfaceMockRecorder) AttachToContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"AttachToContainer",
		reflect.TypeOf((*MockInterface)(nil).AttachToContainer),
		arg0,
		arg1,
		arg2,
	)
}

// ResizeContainerTTY mocks base method
func (m *MockInterface) ResizeContainerTTY(id string, height, width uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeContainerTTY", id, height, width)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResizeContainerTTY indicates an expected call of ResizeContainerTTY
func (mr *MockInterfaceMockRecorder) ResizeContainerTTY(
	id, height, width interface{},
) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ResizeContainerTTY",
		reflect.TypeOf((*MockInterface)(nil).ResizeContainerTTY),
		id,
		height,
		width,
	)
}

// ResizeExecTTY mocks base method
func (m *MockInterface) ResizeExecTTY(id string, height, width uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeExecTTY", id, height, width)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResizeExecTTY indicates an expected call of ResizeExecTTY
func (mr *MockInterfaceMockRecorder) ResizeExecTTY(id, height, width interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"ResizeExecTTY",
		reflect.TypeOf((*MockInterface)(nil).ResizeExecTTY),
		id,
		height,
		width,
	)
}

// GetContainerStats mocks base method
func (m *MockInterface) GetContainerStats(id string) (*types.StatsJSON, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerStats", id)
	ret0, _ := ret[0].(*types.StatsJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerStats indicates an expected call of GetContainerStats
func (mr *MockInterfaceMockRecorder) GetContainerStats(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"GetContainerStats",
		reflect.TypeOf((*MockInterface)(nil).GetContainerStats),
		id,
	)
}
