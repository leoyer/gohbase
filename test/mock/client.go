// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/tsuna/gohbase (interfaces: Client)

package mock

import (
	gomock "github.com/golang/mock/gomock"
	hrpc "github.com/tsuna/gohbase/hrpc"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Append(_param0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Append", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Append(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Append", arg0)
}

func (_m *MockClient) CheckAndPut(_param0 *hrpc.Mutate, _param1 string, _param2 string, _param3 []byte) (bool, error) {
	ret := _m.ctrl.Call(_m, "CheckAndPut", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) CheckAndPut(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckAndPut", arg0, arg1, arg2, arg3)
}

func (_m *MockClient) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockClientRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockClient) Delete(_param0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockClient) Get(_param0 *hrpc.Get) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockClient) Increment(_param0 *hrpc.Mutate) (int64, error) {
	ret := _m.ctrl.Call(_m, "Increment", _param0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Increment(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Increment", arg0)
}

func (_m *MockClient) Put(_param0 *hrpc.Mutate) (*hrpc.Result, error) {
	ret := _m.ctrl.Call(_m, "Put", _param0)
	ret0, _ := ret[0].(*hrpc.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Put(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0)
}

func (_m *MockClient) Scan(_param0 *hrpc.Scan) hrpc.Scanner {
	ret := _m.ctrl.Call(_m, "Scan", _param0)
	ret0, _ := ret[0].(hrpc.Scanner)
	return ret0
}

func (_mr *_MockClientRecorder) Scan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scan", arg0)
}
