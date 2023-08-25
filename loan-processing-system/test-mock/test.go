package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestHelper struct {
	mock.Mock
}

func (th *TestHelper) AssertCalledWithArgs(t *testing.T, methodName string, args ...interface{}) {
	for _, call := range th.Mock.Calls {
		if call.Method == methodName {
			assert.ElementsMatch(t, args, call.Arguments)
		}
	}
}
