package handlers

import (
	"github.com/Adebusy/CCMSService/datastore/sqlserver"
	"github.com/Adebusy/CCMSService/driver"
)

var compService = sqlserver.NewComplaintService(driver.GetDB())
//UserInterface user interface
var UserInterface = sqlserver.NewuserService(driver.GetDB())
func init() {
	//fmt.Println("all handlers")
	UserInterface = sqlserver.NewuserService(driver.GetDB())
	compService = sqlserver.NewComplaintService(driver.GetDB())
}
