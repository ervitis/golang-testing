package users

import (
	"github.com/stretchr/testify/mock"
)

type mocker struct {
	mock.Mock
}

func (m *mocker) ReadData(path string, parser interface{}) error {
	args := m.Called(path, parser)
	return args.Error(0)
}
