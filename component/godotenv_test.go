/**
  @author: honor
  @since: 2021/7/12
  @desc: //TODO
**/
package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnv(t *testing.T) {
	//assert.NotPanics(t, func() {
	//	ReadEnv()
	//})

	assert.NotPanics(t, func() {
		Read2Map()
	})

}
