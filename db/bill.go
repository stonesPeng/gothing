/**
  @author: honor
  @since: 2021/7/7
  @desc: //TODO
**/
package db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx/types"
	"time"
)

const (
	insertSchema = "INSERT INTO bill(channel,charge,removed,version,status) VALUES (?,?,?,?,?)"
	updateSchema = "UPDATE bill SET paid_time=?,status=? where id=?"
	selectSchema = "SELECT * FROM bill WHERE id =?"
)

type Bill struct {
	Id       uint64        `json:"id,omitempty"  db:"id"`
	Channel  uint64        `json:"channel,omitempty" db:"channel"`
	PaidTime time.Time     `json:"paidTime,omitempty" db:"paid_time"`
	Charge   string        `json:"charge,omitempty" db:"charge"`
	Removed  types.BitBool `json:"removed,omitempty" db:"removed"`
	Version  uint32        `json:"version,omitempty" db:"version"`
	Status   uint8         `json:"status,omitempty" db:"status"`
}

func init() {
	DatabaseLink := "root:medtree123456@tcp(192.168.8.94:33061)/test_db"
	dao = &Dao{}
	err := dao.connect(DatabaseLink)
	if err != nil {
		panic(err)
	}
	billRepo = BillRepo{billDao: dao}
	fmt.Println("init repo")
}

var billRepo BillRepo

type BillRepo struct {
	billDao *Dao
}

func (repo BillRepo) CreateBill(channel uint64, charge string) (bill *Bill, er error) {
	bill = &Bill{Channel: channel, Charge: charge, Status: 1}
	id, err := repo.billDao.Create(insertSchema, bill.Channel, bill.Charge, 0, 1, bill.Status)
	if err != nil {
		return nil, errors.New("保存账单失败")
	}
	bill.Id = id
	return bill, nil
}

/**
 * @Description:
 * @receiver repo
 * @param id
 * @param paidTime
 * @param status
 * @return bill
 * @return err
 */
func (repo BillRepo) UpdateBill(id uint64, paidTime time.Time, status uint8) (bill *Bill, err error) {
	if affectedRows, er := repo.billDao.Update(updateSchema, paidTime, status, id); er != nil || affectedRows < 1 {
		return nil, errors.New("更新账单失败")
	} else {
		bills := make([]Bill, 10)
		if err1 := repo.billDao.LoadById(selectSchema, id, &bills); err1 != nil {
			return nil, errors.New("获取账单失败")
		} else {
			return &bills[0], nil
		}
	}
}
