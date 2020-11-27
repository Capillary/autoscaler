/*
Copyright 2017 The Kubernetes Authors.

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

package mocks

import schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
import cloudprovider "k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// NodeGroup is an autogenerated mock type for the NodeGroup type
type NodeGroup struct {
	mock.Mock
}

// Autoprovisioned provides a mock function with given fields:
func (_m *NodeGroup) Autoprovisioned() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *NodeGroup) Create() (cloudprovider.NodeGroup, error) {
	ret := _m.Called()

	var r0 cloudprovider.NodeGroup
	if rf, ok := ret.Get(0).(func() cloudprovider.NodeGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cloudprovider.NodeGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Debug provides a mock function with given fields:
func (_m *NodeGroup) Debug() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DecreaseTargetSize provides a mock function with given fields: delta
func (_m *NodeGroup) DecreaseTargetSize(delta int) error {
	ret := _m.Called(delta)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(delta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields:
func (_m *NodeGroup) Delete() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNodes provides a mock function with given fields: _a0
func (_m *NodeGroup) DeleteNodes(_a0 []*v1.Node) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*v1.Node) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exist provides a mock function with given fields:
func (_m *NodeGroup) Exist() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Id provides a mock function with given fields:
func (_m *NodeGroup) Id() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IncreaseSize provides a mock function with given fields: delta
func (_m *NodeGroup) IncreaseSize(delta int) error {
	ret := _m.Called(delta)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(delta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MaxSize provides a mock function with given fields:
func (_m *NodeGroup) MaxSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MinSize provides a mock function with given fields:
func (_m *NodeGroup) MinSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Nodes provides a mock function with given fields:
func (_m *NodeGroup) Nodes() ([]cloudprovider.Instance, error) {
	ret := _m.Called()

	var r0 []cloudprovider.Instance
	if rf, ok := ret.Get(0).(func() []cloudprovider.Instance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cloudprovider.Instance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TargetSize provides a mock function with given fields:
func (_m *NodeGroup) TargetSize() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TemplateNodeInfo provides a mock function with given fields:
func (_m *NodeGroup) TemplateNodeInfo() (*schedulerframework.NodeInfo, error) {
	ret := _m.Called()

	var r0 *schedulerframework.NodeInfo
	if rf, ok := ret.Get(0).(func() *schedulerframework.NodeInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*schedulerframework.NodeInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
