package sqlserver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"

	md "github.com/Adebusy/CCMSService/models"
)

//IUser user interface
type IUser interface {
	CheckUser(ctx context.Context, RequestBody string) (md.TblUsers, error)
	GetUserByEmail(ctx context.Context, UserEmail string) (md.TblUsers, error)
	//GetUserByUsername(ctx context.Context, UserEmail string) (md.TblUsers, error)
	CreateUser(ctx context.Context, RequestBody md.TblUsers) md.ResponseMessage
	UpdateUser(ctx context.Context, RequestBody md.TblUsers) md.ResponseMessage
	// GetUsers(ctx context.Context) ([]md.TblUsers, error)
	ChangeUserStatusRole(ctx context.Context, RequestBody md.TblUsers) md.ResponseMessage
}

type userService struct {
	DbGorm *gorm.DB
}

//NewuserService interface to be assesible
func NewuserService(DbGorm *gorm.DB) IUser {
	return &userService{DbGorm}
}

func (er userService) CheckUser(ctx context.Context, RequestBody string) (md.TblUsers, error) {
	var UserResp md.TblUsers
	fmt.Println("got here41")
	fmt.Println(RequestBody)
	query := er.DbGorm.Table(`Tbl_users`).Where(`user_name =?`, RequestBody).First(&UserResp)
	if UserResp.ID != 0 {
		fmt.Println("got here4")
		fmt.Println(query)
		return UserResp, nil
	}
	return UserResp, nil
}

var respVal md.ResponseMessage

func (er userService) CreateUser(ctx context.Context, RequestBody md.TblUsers) md.ResponseMessage {
	//var respVal md.ResponseMessage
	fmt.Println("took me here")
	fmt.Println(RequestBody.OfficeLocation)
	err := er.DbGorm.Table(`Tbl_users`).Create(&RequestBody).Error
	fmt.Println("error stage")
	fmt.Println(err)

	if RequestBody.ID > 0 {
		fmt.Println("took me here1")
		respVal.ResponseDescription = strconv.Itoa(RequestBody.ID) + " New user created successfully."
		respVal.ResponseCode = "00"
	} else {
		fmt.Println("took me here2")
		respVal.ResponseDescription = "Error occurred."
		respVal.ResponseCode = "01"
	}
	fmt.Println("took me here3")
	return respVal
}

func (er userService) UpdateUser(ctx context.Context, RequestBody md.TblUsers) md.ResponseMessage {
	var lastn string = RequestBody.LastName
	fmt.Println(lastn)
	//query := er.dbGorm.Model(&RequestBody).Updates(map[string]interface{}{"LastName": lastn, "FirstName": lastn})
	query := er.DbGorm.Table(`Tbl_users`).UpdateColumns(md.TblUsers{LastName: lastn, FirstName: lastn}).Error
	if query != nil {
		respVal.ResponseCode = "01"
		fmt.Println(query.Error)
		respVal.ResponseDescription = "Error occured."
	} else {
		respVal.ResponseCode = "00"
		respVal.ResponseDescription = "Record updated successfully."
	}
	return respVal
}

// func (er userService) GetUsers(ctx context.Context) ([]md.TblUsers, error) {
// 	var respVals []md.TblUsers
// 	query := er.dbGorm.Table(`Tbl_users`).Select(&respVals).Error
// 	if query != nil {
// 		return respVals, query
// 	}
// 	return respVals, nil
// }

func (er userService) ChangeUserStatusRole(ctx context.Context, RequestBody md.TblUsers) md.ResponseMessage {
	query := er.DbGorm.Table(`Tbl_users`).Where(`user_name=?`, RequestBody.UserName).UpdateColumns(md.TblUsers{UserStatus: RequestBody.UserStatus, RoleID: RequestBody.RoleID, UpdatedBy: RequestBody.UpdatedBy}).Error
	if query != nil {
		respVal.ResponseCode = "01"
		fmt.Println(query.Error)
		respVal.ResponseDescription = "Error occurred"
	} else {
		respVal.ResponseCode = "00"
		respVal.ResponseDescription = "Record updated successfully."
	}
	return respVal
}

func (er userService) GetUserByEmail(ctx context.Context, UserEmail string) (md.TblUsers, error) {
	fmt.Println("get here")
	fmt.Println(UserEmail)
	var respVals md.TblUsers
	//er.DbGorm.LogMode(true).Table(`fads`
	query := er.DbGorm.Debug().Table(`Tbl_users`).Where(`email=?`, UserEmail).First(&respVals).Error
	if query != nil {
		return respVals, query
	}
	return respVals, nil
}
