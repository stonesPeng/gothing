/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package go_akka

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.NotPanics(t, func() {
		cfg := new(Config).parseConfig("./cfg.yaml")
		fmt.Printf(" port=%d,debug=%t", cfg.Port, cfg.Debug)
	})
}
