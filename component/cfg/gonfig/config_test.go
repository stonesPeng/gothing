/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package gonfig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.NotPanics(t, func() {
		fmt.Println("Dev Conf :")
		cfg := ParseConfig()
		fmt.Printf("userName=%s,password=%s,port=%s,host=%s,name=%s \n", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_PORT, cfg.DB_HOST, cfg.DB_NAME)
		cfg = ParseConfig("prd")
		fmt.Printf("userName=%s,password=%s,port=%s,host=%s,name=%s \n", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_PORT, cfg.DB_HOST, cfg.DB_NAME)
	})

}
