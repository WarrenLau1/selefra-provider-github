package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-github/constants"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

type MockRepositoriesService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesServiceMockRecorder
}

type MockRepositoriesServiceMockRecorder struct {
	mock *MockRepositoriesService
}

func NewMockRepositoriesService(ctrl *gomock.Controller) *MockRepositoriesService {
	mock := &MockRepositoriesService{ctrl: ctrl}
	mock.recorder = &MockRepositoriesServiceMockRecorder{mock}
	return mock
}

func (m *MockRepositoriesService) EXPECT() *MockRepositoriesServiceMockRecorder {
	return m.recorder
}

func (m *MockRepositoriesService) ListByOrg(arg0 context.Context, arg1 string, arg2 *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListByOrg, arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockRepositoriesServiceMockRecorder) ListByOrg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListByOrg, reflect.TypeOf((*MockRepositoriesService)(nil).ListByOrg), arg0, arg1, arg2)
}
