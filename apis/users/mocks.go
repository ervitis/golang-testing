package users

import (
	"encoding/json"
	"fmt"
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

func mockUsers(n ...int) []*User {
	var users []*User
	var end int

	if n == nil {
		end = 150
	} else {
		end = n[0]
	}

	for i := 1; i < end; i++ {
		user := &User{Surname: nil, Name: fmt.Sprintf("test%d", i), Id: i, Gender: nil, Email: "test@test.com", Country: "Spain"}
		users = append(users, user)
	}
	return users
}
