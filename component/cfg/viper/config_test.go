/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package viper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.NotPanics(t, func() {
		readConfig()
	})
}
