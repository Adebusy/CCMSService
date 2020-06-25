package sqlserver

import (
	"context"
	"fmt"

	modules "github.com/Adebusy/CCMSService/models"
	"github.com/jinzhu/gorm"
)

//var connect = con.NewClient(dbGorm)

//var dbGorm *gorm.DB
//var errGorm error

type complaintService struct {
	dbGorm *gorm.DB
}

//NewComplaintService complaint
func NewComplaintService(db *gorm.DB) IComplaint {
	return &complaintService{db}
}

//IComplaint interface
type IComplaint interface {
	LogComplaint(ctx context.Context, RequestBody modules.TblCases) (int, error)
	//GetSingleComplaintByComplaintID(ctx context.Context, RequestBody modules.TblCases) (int error)
}

func (ts complaintService) LogComplaint(ctx context.Context, RequestBody modules.TblCases) (int, error) {
	fmt.Println("got to LogComplaint")
	query := ts.dbGorm.Table(`tbl_cases`).Create(&RequestBody).Error
	if query != nil {
		fmt.Println("got to LogComplaint 1")
		return RequestBody.ID, query
	}
	return RequestBody.ID, nil
}
