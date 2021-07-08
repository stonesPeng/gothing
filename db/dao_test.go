/**
  @author: honor
  @since: 2021/7/7
  @desc: //TODO
**/
package db

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/**
 * @Description:  测试数据库
 * @param t
 */
func TestDB(t *testing.T) {
	fmt.Println("into  testing")
	uid := uint64(0)
	assert.NotPanics(t, func() {
		bill, err := billRepo.CreateBill(uint64(10), "123.123")
		if err != nil {
			panic(err)
		}
		json, _ := json.Marshal(bill)
		uid = bill.Id
		fmt.Printf("bill is %s", json)
	})

	assert.NotPanics(t, func() {
		bill, err := billRepo.UpdateBill(uid, time.Now(), 2)
		if err != nil {
			panic(err)
		}
		json, _ := json.Marshal(bill)
		uid = bill.Id
		fmt.Printf("bill is %s", json)
	})
}
