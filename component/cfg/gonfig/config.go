/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package gonfig

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func ParseConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fName := fmt.Sprintf("./%s_config.json", env)
	if err := gonfig.GetConf(fName, &configuration); err != nil {
		panic(err)
	}
	return configuration
}
