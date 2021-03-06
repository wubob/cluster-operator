// Code generated by MockGen. DO NOT EDIT.
// Source: ./jobsyncstrategy.go

// Package controller is a generated GoMock package.
package controller

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
)

// MockJobSyncStrategy is a mock of JobSyncStrategy interface
type MockJobSyncStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockJobSyncStrategyMockRecorder
}

// MockJobSyncStrategyMockRecorder is the mock recorder for MockJobSyncStrategy
type MockJobSyncStrategyMockRecorder struct {
	mock *MockJobSyncStrategy
}

// NewMockJobSyncStrategy creates a new mock instance
func NewMockJobSyncStrategy(ctrl *gomock.Controller) *MockJobSyncStrategy {
	mock := &MockJobSyncStrategy{ctrl: ctrl}
	mock.recorder = &MockJobSyncStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJobSyncStrategy) EXPECT() *MockJobSyncStrategyMockRecorder {
	return m.recorder
}

// GetOwner mocks base method
func (m *MockJobSyncStrategy) GetOwner(key string) (v10.Object, error) {
	ret := m.ctrl.Call(m, "GetOwner", key)
	ret0, _ := ret[0].(v10.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwner indicates an expected call of GetOwner
func (mr *MockJobSyncStrategyMockRecorder) GetOwner(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwner", reflect.TypeOf((*MockJobSyncStrategy)(nil).GetOwner), key)
}

// DoesOwnerNeedProcessing mocks base method
func (m *MockJobSyncStrategy) DoesOwnerNeedProcessing(owner v10.Object) bool {
	ret := m.ctrl.Call(m, "DoesOwnerNeedProcessing", owner)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DoesOwnerNeedProcessing indicates an expected call of DoesOwnerNeedProcessing
func (mr *MockJobSyncStrategyMockRecorder) DoesOwnerNeedProcessing(owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesOwnerNeedProcessing", reflect.TypeOf((*MockJobSyncStrategy)(nil).DoesOwnerNeedProcessing), owner)
}

// GetJobFactory mocks base method
func (m *MockJobSyncStrategy) GetJobFactory(owner v10.Object, deleting bool) (JobFactory, error) {
	ret := m.ctrl.Call(m, "GetJobFactory", owner, deleting)
	ret0, _ := ret[0].(JobFactory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobFactory indicates an expected call of GetJobFactory
func (mr *MockJobSyncStrategyMockRecorder) GetJobFactory(owner, deleting interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobFactory", reflect.TypeOf((*MockJobSyncStrategy)(nil).GetJobFactory), owner, deleting)
}

// DeepCopyOwner mocks base method
func (m *MockJobSyncStrategy) DeepCopyOwner(owner v10.Object) v10.Object {
	ret := m.ctrl.Call(m, "DeepCopyOwner", owner)
	ret0, _ := ret[0].(v10.Object)
	return ret0
}

// DeepCopyOwner indicates an expected call of DeepCopyOwner
func (mr *MockJobSyncStrategyMockRecorder) DeepCopyOwner(owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeepCopyOwner", reflect.TypeOf((*MockJobSyncStrategy)(nil).DeepCopyOwner), owner)
}

// SetOwnerJobSyncCondition mocks base method
func (m *MockJobSyncStrategy) SetOwnerJobSyncCondition(owner v10.Object, conditionType JobSyncConditionType, status v1.ConditionStatus, reason, message string, updateConditionCheck UpdateConditionCheck) {
	m.ctrl.Call(m, "SetOwnerJobSyncCondition", owner, conditionType, status, reason, message, updateConditionCheck)
}

// SetOwnerJobSyncCondition indicates an expected call of SetOwnerJobSyncCondition
func (mr *MockJobSyncStrategyMockRecorder) SetOwnerJobSyncCondition(owner, conditionType, status, reason, message, updateConditionCheck interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwnerJobSyncCondition", reflect.TypeOf((*MockJobSyncStrategy)(nil).SetOwnerJobSyncCondition), owner, conditionType, status, reason, message, updateConditionCheck)
}

// OnJobCompletion mocks base method
func (m *MockJobSyncStrategy) OnJobCompletion(owner v10.Object) {
	m.ctrl.Call(m, "OnJobCompletion", owner)
}

// OnJobCompletion indicates an expected call of OnJobCompletion
func (mr *MockJobSyncStrategyMockRecorder) OnJobCompletion(owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnJobCompletion", reflect.TypeOf((*MockJobSyncStrategy)(nil).OnJobCompletion), owner)
}

// UpdateOwnerStatus mocks base method
func (m *MockJobSyncStrategy) UpdateOwnerStatus(original, owner v10.Object) error {
	ret := m.ctrl.Call(m, "UpdateOwnerStatus", original, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOwnerStatus indicates an expected call of UpdateOwnerStatus
func (mr *MockJobSyncStrategyMockRecorder) UpdateOwnerStatus(original, owner interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOwnerStatus", reflect.TypeOf((*MockJobSyncStrategy)(nil).UpdateOwnerStatus), original, owner)
}
