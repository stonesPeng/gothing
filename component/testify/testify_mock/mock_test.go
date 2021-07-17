/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package testify_mock

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type dbMock struct {
	mock.Mock
}

func (m *dbMock) FetchMessage(lang string) (bool, error) {
	args := m.Called(lang)
	return args.Bool(0), args.Error(1)
}

func (m *dbMock) FetchDefaultMessage() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func TestMockMethodWithoutArgs(t *testing.T) {
	theDBMock := dbMock{}
	theDBMock.On("FetchDefaultMessage").Return("foofofofof", nil)
	g := greeter{theDBMock, "en"}
	assert.Equal(t, "Message is: foofofofof", g.Greet())
}
