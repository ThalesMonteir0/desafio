// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/telemetry.go
//
// Generated by this command:
//
//	mockgen -source=internal/ports/telemetry.go -destination=internal/test/mocks/telemetry.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/ThalesMonteir0/desafio/internal/domain"
	err_rest "github.com/ThalesMonteir0/desafio/pkg/err_rest"
	gomock "go.uber.org/mock/gomock"
)

// MockTelemetryRepository is a mock of TelemetryRepository interface.
type MockTelemetryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTelemetryRepositoryMockRecorder
}

// MockTelemetryRepositoryMockRecorder is the mock recorder for MockTelemetryRepository.
type MockTelemetryRepositoryMockRecorder struct {
	mock *MockTelemetryRepository
}

// NewMockTelemetryRepository creates a new mock instance.
func NewMockTelemetryRepository(ctrl *gomock.Controller) *MockTelemetryRepository {
	mock := &MockTelemetryRepository{ctrl: ctrl}
	mock.recorder = &MockTelemetryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemetryRepository) EXPECT() *MockTelemetryRepositoryMockRecorder {
	return m.recorder
}

// CreateGps mocks base method.
func (m *MockTelemetryRepository) CreateGps(gpsDomain domain.GpsDomain) *err_rest.ErrRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGps", gpsDomain)
	ret0, _ := ret[0].(*err_rest.ErrRest)
	return ret0
}

// CreateGps indicates an expected call of CreateGps.
func (mr *MockTelemetryRepositoryMockRecorder) CreateGps(gpsDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGps", reflect.TypeOf((*MockTelemetryRepository)(nil).CreateGps), gpsDomain)
}

// CreateGyroscope mocks base method.
func (m *MockTelemetryRepository) CreateGyroscope(gyroscopeDomain domain.GyroscopeDomain) *err_rest.ErrRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGyroscope", gyroscopeDomain)
	ret0, _ := ret[0].(*err_rest.ErrRest)
	return ret0
}

// CreateGyroscope indicates an expected call of CreateGyroscope.
func (mr *MockTelemetryRepositoryMockRecorder) CreateGyroscope(gyroscopeDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGyroscope", reflect.TypeOf((*MockTelemetryRepository)(nil).CreateGyroscope), gyroscopeDomain)
}

// CreatePhoto mocks base method.
func (m *MockTelemetryRepository) CreatePhoto(photoDomain domain.PhotoDomain) *err_rest.ErrRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoto", photoDomain)
	ret0, _ := ret[0].(*err_rest.ErrRest)
	return ret0
}

// CreatePhoto indicates an expected call of CreatePhoto.
func (mr *MockTelemetryRepositoryMockRecorder) CreatePhoto(photoDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoto", reflect.TypeOf((*MockTelemetryRepository)(nil).CreatePhoto), photoDomain)
}

// FindDeviceByMAC mocks base method.
func (m *MockTelemetryRepository) FindDeviceByMAC(macAddress string) (domain.DeviceDomain, *err_rest.ErrRest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindDeviceByMAC", macAddress)
	ret0, _ := ret[0].(domain.DeviceDomain)
	ret1, _ := ret[1].(*err_rest.ErrRest)
	return ret0, ret1
}

// FindDeviceByMAC indicates an expected call of FindDeviceByMAC.
func (mr *MockTelemetryRepositoryMockRecorder) FindDeviceByMAC(macAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindDeviceByMAC", reflect.TypeOf((*MockTelemetryRepository)(nil).FindDeviceByMAC), macAddress)
}

// MockTelemetryService is a mock of TelemetryService interface.
type MockTelemetryService struct {
	ctrl     *gomock.Controller
	recorder *MockTelemetryServiceMockRecorder
}

// MockTelemetryServiceMockRecorder is the mock recorder for MockTelemetryService.
type MockTelemetryServiceMockRecorder struct {
	mock *MockTelemetryService
}

// NewMockTelemetryService creates a new mock instance.
func NewMockTelemetryService(ctrl *gomock.Controller) *MockTelemetryService {
	mock := &MockTelemetryService{ctrl: ctrl}
	mock.recorder = &MockTelemetryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemetryService) EXPECT() *MockTelemetryServiceMockRecorder {
	return m.recorder
}

// CreateGpsService mocks base method.
func (m *MockTelemetryService) CreateGpsService(gpsDomain domain.GpsDomain) *err_rest.ErrRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGpsService", gpsDomain)
	ret0, _ := ret[0].(*err_rest.ErrRest)
	return ret0
}

// CreateGpsService indicates an expected call of CreateGpsService.
func (mr *MockTelemetryServiceMockRecorder) CreateGpsService(gpsDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGpsService", reflect.TypeOf((*MockTelemetryService)(nil).CreateGpsService), gpsDomain)
}

// CreateGyroscopeService mocks base method.
func (m *MockTelemetryService) CreateGyroscopeService(gyroscopeDomain domain.GyroscopeDomain) *err_rest.ErrRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGyroscopeService", gyroscopeDomain)
	ret0, _ := ret[0].(*err_rest.ErrRest)
	return ret0
}

// CreateGyroscopeService indicates an expected call of CreateGyroscopeService.
func (mr *MockTelemetryServiceMockRecorder) CreateGyroscopeService(gyroscopeDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGyroscopeService", reflect.TypeOf((*MockTelemetryService)(nil).CreateGyroscopeService), gyroscopeDomain)
}

// CreatePhotoService mocks base method.
func (m *MockTelemetryService) CreatePhotoService(photoDomain domain.PhotoDomain) *err_rest.ErrRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhotoService", photoDomain)
	ret0, _ := ret[0].(*err_rest.ErrRest)
	return ret0
}

// CreatePhotoService indicates an expected call of CreatePhotoService.
func (mr *MockTelemetryServiceMockRecorder) CreatePhotoService(photoDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhotoService", reflect.TypeOf((*MockTelemetryService)(nil).CreatePhotoService), photoDomain)
}
