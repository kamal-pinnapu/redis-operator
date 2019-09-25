// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetNumberSentinelSlavesInMemory provides a mock function with given fields: ip
func (_m *Client) GetNumberSentinelSlavesInMemory(ip string) (int32, error) {
	ret := _m.Called(ip)

	var r0 int32
	if rf, ok := ret.Get(0).(func(string) int32); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberSentinelsInMemory provides a mock function with given fields: ip
func (_m *Client) GetNumberSentinelsInMemory(ip string) (int32, error) {
	ret := _m.Called(ip)

	var r0 int32
	if rf, ok := ret.Get(0).(func(string) int32); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSentinelMonitor provides a mock function with given fields: ip
func (_m *Client) GetSentinelMonitor(ip string) (string, error) {
	ret := _m.Called(ip)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSlaveOf provides a mock function with given fields: ip
func (_m *Client) GetSlaveOf(ip string) (string, error) {
	ret := _m.Called(ip)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsMaster provides a mock function with given fields: ip
func (_m *Client) IsMaster(ip string) (bool, error) {
	ret := _m.Called(ip)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeMaster provides a mock function with given fields: ip
func (_m *Client) MakeMaster(ip string) error {
	ret := _m.Called(ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MakeSlaveOf provides a mock function with given fields: ip, masterIP
func (_m *Client) MakeSlaveOf(ip string, masterIP string) error {
	ret := _m.Called(ip, masterIP)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(ip, masterIP)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MonitorRedis provides a mock function with given fields: ip, monitor, quorum
func (_m *Client) MonitorRedis(ip string, monitor string, quorum string) error {
	ret := _m.Called(ip, monitor, quorum)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(ip, monitor, quorum)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetSentinel provides a mock function with given fields: ip
func (_m *Client) ResetSentinel(ip string) error {
	ret := _m.Called(ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCustomRedisConfig provides a mock function with given fields: ip, configs
func (_m *Client) SetCustomRedisConfig(ip string, configs []string) error {
	ret := _m.Called(ip, configs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(ip, configs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCustomSentinelConfig provides a mock function with given fields: ip, configs
func (_m *Client) SetCustomSentinelConfig(ip string, configs []string) error {
	ret := _m.Called(ip, configs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(ip, configs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCustomSentinelConfig provides a mock function with given fields: password
func (_m *Client) SetRedisAuth(password string) {
	_m.Called(password)
}

func (_m *Client) GetRedisAuth() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ""
	}

	return r0
}
