/**
  @author: honor
  @since: 2021/7/12
  @desc: //TODO
**/
package syntax

import (
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	n := foo()
	fmt.Println("received ", n)
}
