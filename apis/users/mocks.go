package users

import (
	"github.com/stretchr/testify/mock"
)

type mockGetUsers struct {
	mock.Mock
}

func (m *mockGetUsers) ReadData(path string, parser interface{}) error {
	args := m.Called(path, parser)
	return args.Error(0)
}
