package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type JccCustomer struct {
	gorm.Model
	ID                int64   `gorm:"column:id"`
	Name              string  `gorm:"column:name"`
	Address           string  `gorm:"column:address"`
	Phone             string  `gorm:"column:phone"`
	IsDel             int     `gorm:"column:is_del"`
	CompanyId         int     `gorm:"column:company_id"`
	Balance           float64 `gorm:"column:balance"`
	Overdraft         float64 `gorm:"column:overdraft"`
	GiveBalance       float64 `gorm:"column:giveBalance"`
	AccruingTopUp     float64 `gorm:"column:accruing_top_up"`
	AccumulatedOutlay float64 `gorm:"column:accumulated_outlay"`
	DeletedAt         int     `gorm:"default:'0'"`
}

func init() {

}

func GetCustomer() {
	var Customer JccCustomer
	//Db.First(&Customer)
	Db.First(&Customer)
	fmt.Println(fmt.Sprintf("%+v", Customer))
	//fmt.Println(Customer)
	defer Db.Close()
}
