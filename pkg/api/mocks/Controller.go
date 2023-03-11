// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import events "github.com/aws/aws-lambda-go/events"
import mock "github.com/stretchr/testify/mock"

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, req
func (_m *Controller) Call(ctx context.Context, req *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 events.APIGatewayProxyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *events.APIGatewayProxyRequest) events.APIGatewayProxyResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(events.APIGatewayProxyResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *events.APIGatewayProxyRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
