package users

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
)

type mocker struct {
	mock.Mock
}

func (m *mocker) ReadData(path string) ([]byte, error) {
	args := m.Called(path)
	b, _ := json.Marshal(args.Get(0))
	return b, args.Error(1)
}

func mockUsers() []byte {
	return []byte(`[{"id": 1, "name": "test", "surname": null, "email": "test@test.com", "gender": null, "country": "Spain"}]`)
}
