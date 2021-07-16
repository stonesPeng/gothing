/**
  @author: honor
  @since: 2021/7/14
  @desc: //TODO
**/
package basic

import "fmt"

func slice_handle() {
	arr := []string{"1", "2", "3", "4", "5"}
	s1 := arr[1:2]
	s2 := arr[2:3]
	for _, i := range s1 {
		fmt.Println(i)
	}
	for _, i := range s2 {
		fmt.Println(i)
	}
}
