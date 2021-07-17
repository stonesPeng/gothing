/**
  @author: honor
  @since: 2021/7/12
  @desc: //TODO
**/
package godotenv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnv(t *testing.T) {
	//testify.NotPanics(t, func() {
	//	ReadEnv()
	//})

	assert.NotPanics(t, func() {
		Read2Map()
	})

}
