/**
  @author: honor
  @since: 2021/7/16
  @desc: //TODO
**/
package go_akka

import "github.com/go-akka/configuration"

type Config struct {
	Port  uint16 `yaml:"port"`
	Debug bool   `yaml:"debug"`
	Raw   *configuration.Config
}

/**
 * @Description:
 * @receiver c
 * @param path
 * @return *Config
 */
func (c *Config) parseConfig(path string) *Config {
	c.Raw = configuration.LoadConfig(path)
	c.Port = uint16(c.Raw.GetInt32("port", 8008))
	c.Debug = c.Raw.GetBoolean("debug", false)
	return c
}
