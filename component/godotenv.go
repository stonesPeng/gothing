/**
  @author: honor
  @since: 2021/7/12
  @desc: //读取配置，存储到os环境变量
**/
package component

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ReadEnv() {
	value1 := os.Getenv("Key1")
	value2 := os.Getenv("Key2")
	value3 := os.Getenv("Key3")
	if value1 == "" || value2 == "" {
		panic("load error")
	}
	fmt.Printf("key1=%s,key2=%s，key3=%s", value1, value2, value3)
}
