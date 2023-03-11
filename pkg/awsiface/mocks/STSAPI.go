// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	request "github.com/aws/aws-sdk-go/aws/request"
	mock "github.com/stretchr/testify/mock"

	sts "github.com/aws/aws-sdk-go/service/sts"
)

// STSAPI is an autogenerated mock type for the STSAPI type
type STSAPI struct {
	mock.Mock
}

// AssumeRole provides a mock function with given fields: _a0
func (_m *STSAPI) AssumeRole(_a0 *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.AssumeRoleOutput
	if rf, ok := ret.Get(0).(func(*sts.AssumeRoleInput) *sts.AssumeRoleOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.AssumeRoleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.AssumeRoleInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssumeRoleRequest provides a mock function with given fields: _a0
func (_m *STSAPI) AssumeRoleRequest(_a0 *sts.AssumeRoleInput) (*request.Request, *sts.AssumeRoleOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.AssumeRoleInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.AssumeRoleOutput
	if rf, ok := ret.Get(1).(func(*sts.AssumeRoleInput) *sts.AssumeRoleOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.AssumeRoleOutput)
		}
	}

	return r0, r1
}

// AssumeRoleWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) AssumeRoleWithContext(_a0 context.Context, _a1 *sts.AssumeRoleInput, _a2 ...request.Option) (*sts.AssumeRoleOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.AssumeRoleOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.AssumeRoleInput, ...request.Option) *sts.AssumeRoleOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.AssumeRoleOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.AssumeRoleInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssumeRoleWithSAML provides a mock function with given fields: _a0
func (_m *STSAPI) AssumeRoleWithSAML(_a0 *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.AssumeRoleWithSAMLOutput
	if rf, ok := ret.Get(0).(func(*sts.AssumeRoleWithSAMLInput) *sts.AssumeRoleWithSAMLOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.AssumeRoleWithSAMLOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.AssumeRoleWithSAMLInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssumeRoleWithSAMLRequest provides a mock function with given fields: _a0
func (_m *STSAPI) AssumeRoleWithSAMLRequest(_a0 *sts.AssumeRoleWithSAMLInput) (*request.Request, *sts.AssumeRoleWithSAMLOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.AssumeRoleWithSAMLInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.AssumeRoleWithSAMLOutput
	if rf, ok := ret.Get(1).(func(*sts.AssumeRoleWithSAMLInput) *sts.AssumeRoleWithSAMLOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.AssumeRoleWithSAMLOutput)
		}
	}

	return r0, r1
}

// AssumeRoleWithSAMLWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) AssumeRoleWithSAMLWithContext(_a0 context.Context, _a1 *sts.AssumeRoleWithSAMLInput, _a2 ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.AssumeRoleWithSAMLOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.AssumeRoleWithSAMLInput, ...request.Option) *sts.AssumeRoleWithSAMLOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.AssumeRoleWithSAMLOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.AssumeRoleWithSAMLInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssumeRoleWithWebIdentity provides a mock function with given fields: _a0
func (_m *STSAPI) AssumeRoleWithWebIdentity(_a0 *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.AssumeRoleWithWebIdentityOutput
	if rf, ok := ret.Get(0).(func(*sts.AssumeRoleWithWebIdentityInput) *sts.AssumeRoleWithWebIdentityOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.AssumeRoleWithWebIdentityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.AssumeRoleWithWebIdentityInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssumeRoleWithWebIdentityRequest provides a mock function with given fields: _a0
func (_m *STSAPI) AssumeRoleWithWebIdentityRequest(_a0 *sts.AssumeRoleWithWebIdentityInput) (*request.Request, *sts.AssumeRoleWithWebIdentityOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.AssumeRoleWithWebIdentityInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.AssumeRoleWithWebIdentityOutput
	if rf, ok := ret.Get(1).(func(*sts.AssumeRoleWithWebIdentityInput) *sts.AssumeRoleWithWebIdentityOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.AssumeRoleWithWebIdentityOutput)
		}
	}

	return r0, r1
}

// AssumeRoleWithWebIdentityWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) AssumeRoleWithWebIdentityWithContext(_a0 context.Context, _a1 *sts.AssumeRoleWithWebIdentityInput, _a2 ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.AssumeRoleWithWebIdentityOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.AssumeRoleWithWebIdentityInput, ...request.Option) *sts.AssumeRoleWithWebIdentityOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.AssumeRoleWithWebIdentityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.AssumeRoleWithWebIdentityInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeAuthorizationMessage provides a mock function with given fields: _a0
func (_m *STSAPI) DecodeAuthorizationMessage(_a0 *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.DecodeAuthorizationMessageOutput
	if rf, ok := ret.Get(0).(func(*sts.DecodeAuthorizationMessageInput) *sts.DecodeAuthorizationMessageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.DecodeAuthorizationMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.DecodeAuthorizationMessageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeAuthorizationMessageRequest provides a mock function with given fields: _a0
func (_m *STSAPI) DecodeAuthorizationMessageRequest(_a0 *sts.DecodeAuthorizationMessageInput) (*request.Request, *sts.DecodeAuthorizationMessageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.DecodeAuthorizationMessageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.DecodeAuthorizationMessageOutput
	if rf, ok := ret.Get(1).(func(*sts.DecodeAuthorizationMessageInput) *sts.DecodeAuthorizationMessageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.DecodeAuthorizationMessageOutput)
		}
	}

	return r0, r1
}

// DecodeAuthorizationMessageWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) DecodeAuthorizationMessageWithContext(_a0 context.Context, _a1 *sts.DecodeAuthorizationMessageInput, _a2 ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.DecodeAuthorizationMessageOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.DecodeAuthorizationMessageInput, ...request.Option) *sts.DecodeAuthorizationMessageOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.DecodeAuthorizationMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.DecodeAuthorizationMessageInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccessKeyInfo provides a mock function with given fields: _a0
func (_m *STSAPI) GetAccessKeyInfo(_a0 *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.GetAccessKeyInfoOutput
	if rf, ok := ret.Get(0).(func(*sts.GetAccessKeyInfoInput) *sts.GetAccessKeyInfoOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetAccessKeyInfoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.GetAccessKeyInfoInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccessKeyInfoRequest provides a mock function with given fields: _a0
func (_m *STSAPI) GetAccessKeyInfoRequest(_a0 *sts.GetAccessKeyInfoInput) (*request.Request, *sts.GetAccessKeyInfoOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.GetAccessKeyInfoInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.GetAccessKeyInfoOutput
	if rf, ok := ret.Get(1).(func(*sts.GetAccessKeyInfoInput) *sts.GetAccessKeyInfoOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.GetAccessKeyInfoOutput)
		}
	}

	return r0, r1
}

// GetAccessKeyInfoWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) GetAccessKeyInfoWithContext(_a0 context.Context, _a1 *sts.GetAccessKeyInfoInput, _a2 ...request.Option) (*sts.GetAccessKeyInfoOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.GetAccessKeyInfoOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.GetAccessKeyInfoInput, ...request.Option) *sts.GetAccessKeyInfoOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetAccessKeyInfoOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.GetAccessKeyInfoInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCallerIdentity provides a mock function with given fields: _a0
func (_m *STSAPI) GetCallerIdentity(_a0 *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.GetCallerIdentityOutput
	if rf, ok := ret.Get(0).(func(*sts.GetCallerIdentityInput) *sts.GetCallerIdentityOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetCallerIdentityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.GetCallerIdentityInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCallerIdentityRequest provides a mock function with given fields: _a0
func (_m *STSAPI) GetCallerIdentityRequest(_a0 *sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.GetCallerIdentityInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.GetCallerIdentityOutput
	if rf, ok := ret.Get(1).(func(*sts.GetCallerIdentityInput) *sts.GetCallerIdentityOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.GetCallerIdentityOutput)
		}
	}

	return r0, r1
}

// GetCallerIdentityWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) GetCallerIdentityWithContext(_a0 context.Context, _a1 *sts.GetCallerIdentityInput, _a2 ...request.Option) (*sts.GetCallerIdentityOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.GetCallerIdentityOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.GetCallerIdentityInput, ...request.Option) *sts.GetCallerIdentityOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetCallerIdentityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.GetCallerIdentityInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFederationToken provides a mock function with given fields: _a0
func (_m *STSAPI) GetFederationToken(_a0 *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.GetFederationTokenOutput
	if rf, ok := ret.Get(0).(func(*sts.GetFederationTokenInput) *sts.GetFederationTokenOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetFederationTokenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.GetFederationTokenInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFederationTokenRequest provides a mock function with given fields: _a0
func (_m *STSAPI) GetFederationTokenRequest(_a0 *sts.GetFederationTokenInput) (*request.Request, *sts.GetFederationTokenOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.GetFederationTokenInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.GetFederationTokenOutput
	if rf, ok := ret.Get(1).(func(*sts.GetFederationTokenInput) *sts.GetFederationTokenOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.GetFederationTokenOutput)
		}
	}

	return r0, r1
}

// GetFederationTokenWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) GetFederationTokenWithContext(_a0 context.Context, _a1 *sts.GetFederationTokenInput, _a2 ...request.Option) (*sts.GetFederationTokenOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.GetFederationTokenOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.GetFederationTokenInput, ...request.Option) *sts.GetFederationTokenOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetFederationTokenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.GetFederationTokenInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionToken provides a mock function with given fields: _a0
func (_m *STSAPI) GetSessionToken(_a0 *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sts.GetSessionTokenOutput
	if rf, ok := ret.Get(0).(func(*sts.GetSessionTokenInput) *sts.GetSessionTokenOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetSessionTokenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sts.GetSessionTokenInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSessionTokenRequest provides a mock function with given fields: _a0
func (_m *STSAPI) GetSessionTokenRequest(_a0 *sts.GetSessionTokenInput) (*request.Request, *sts.GetSessionTokenOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sts.GetSessionTokenInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sts.GetSessionTokenOutput
	if rf, ok := ret.Get(1).(func(*sts.GetSessionTokenInput) *sts.GetSessionTokenOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sts.GetSessionTokenOutput)
		}
	}

	return r0, r1
}

// GetSessionTokenWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *STSAPI) GetSessionTokenWithContext(_a0 context.Context, _a1 *sts.GetSessionTokenInput, _a2 ...request.Option) (*sts.GetSessionTokenOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *sts.GetSessionTokenOutput
	if rf, ok := ret.Get(0).(func(context.Context, *sts.GetSessionTokenInput, ...request.Option) *sts.GetSessionTokenOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sts.GetSessionTokenOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sts.GetSessionTokenInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
