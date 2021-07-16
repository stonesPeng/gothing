/**
  @author: honor
  @since: 2021/7/12
  @desc: //TODO
**/
package syntax

import "fmt"

func foo() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	m := 1
	panic("foo:fail")
	m = 2
	return m

}
