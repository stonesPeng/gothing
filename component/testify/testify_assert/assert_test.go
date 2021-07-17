/**
  @author: honor
  @since: 2021/7/16
  @desc: //断言
**/
package testify_assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	//如果有多次断言，使用NEW这种方式
	assert := assert.New(t)
	assert.Equal(123, 123, "they should be equal")
	assert.NotEqual(123, 456, "they should not be equal")
	assert.Nil(nil)
	object := Object{name: "123"}
	if assert.NotNil(object) {
		assert.Equal("123", object.name)
	}
}

type Object struct {
	name string
}
