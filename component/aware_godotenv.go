/**
  @author: honor
  @since: 2021/7/12
  @desc: //TODO
**/
package component

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func AwareReadOs() {
	value1 := os.Getenv("Key1")
	value2 := os.Getenv("Key2")
	value3 := os.Getenv("Key3")
	if value1 == "" || value2 == "" {
		panic("load error")
	}
	fmt.Printf("key1=%s,key2=%s，key3=%s", value1, value2, value3)
}

func Read2Map() {
	maps := map[string]string{}
	maps, err := godotenv.Read()

	if err != nil {
		return
	}

	fmt.Printf("key1=%s,key2=%s，key3=%s", maps["Key1"], maps["Key2"], maps["Key3"])
}
