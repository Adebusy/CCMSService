package sqlserver

import (
	"context"
	"fmt"

	modules "github.com/Adebusy/CCMSService/models"
	"github.com/jinzhu/gorm"
)

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
	GetSingleComplaintByComplaintID(ctx context.Context, RequestBody string) (modules.TblCases, error)
	CreateCompliantCategory(ctx context.Context, RequestBody modules.TblComplaintCategories) (int, error)
	CheckCompliantCategory(ctx context.Context, RequestBody modules.ComplaintCategories) int
	FetchCompliantCategories(ctx context.Context) []modules.TblComplaintCategories
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

func (ts complaintService) GetSingleComplaintByComplaintID(ctx context.Context, RequestBody string) (modules.TblCases, error) {
	var mycase modules.TblCases
	query := ts.dbGorm.Table("tbl_cases").Where("ref_no=?", RequestBody).First(&mycase).Error
	if query != nil {
		return mycase, query
	}
	return mycase, nil
}

func (ts complaintService) CreateCompliantCategory(ctx context.Context, RequestBody modules.TblComplaintCategories) (int, error) {
	query := ts.dbGorm.Table("Tbl_Complaint_Categories").Create(&RequestBody).Error
	if query != nil {
		return 0, query
	}
	return RequestBody.ID, nil
}

func (ts complaintService) CheckCompliantCategory(ctx context.Context, RequestBody modules.ComplaintCategories) int {
	var mycase modules.TblComplaintCategories
	ts.dbGorm.Table("Tbl_Complaint_Categories").Where("Category=?", RequestBody.Category).First(&mycase)
	return mycase.ID
}
func (ts complaintService) FetchCompliantCategories(ctx context.Context) []modules.TblComplaintCategories {
	var mycase []modules.TblComplaintCategories
	ts.dbGorm.Table("Tbl_Complaint_Categories").Find(&mycase)
	return mycase
}
