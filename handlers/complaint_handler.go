package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Adebusy/CCMSService/utilities"

	md "github.com/Adebusy/CCMSService/models"
	"github.com/gin-gonic/gin"
)

//var compService = sqlserver.NewComplaintService(driver.GetDB())
//UserInterface user interface
//var UserInterface = sqlserver.NewuserService(driver.GetDB())

//func init() {
//	UserInterface = sqlserver.NewuserService(driver.GetDB())
//	compService = sqlserver.NewComplaintService(driver.GetDB())
//}
// LogComplaintRequest godoc
// @Summary create new compliant
// @Produce json
// @Param user body models.Cases true "Create new complaint"
// @Success 200 {object} models.ResponseMessage
// @Router /complaint/LogComplaintRequest/ [post]
func LogComplaintRequest(ctx *gin.Context) {
	fmt.Println("got to here cases 1")
	var NewCase md.Cases
	var tblcase md.TblCases

	err := ctx.ShouldBindJSON(&NewCase)
	if err != nil {
		resp.ResponseDescription = err.Error()
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("got to here cases 21")
	if NewCase.AddedBy == "" {
		resp.ResponseDescription = "Initiators name is required for this request to be treated."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("got to here cases 21")
	//confirm initiators name
	queryInitiator, _ := UserInterface.CheckUser(ctx, NewCase.AddedBy)
	if queryInitiator.ID <= 0 {
		resp.ResponseDescription = "Initiators does not exist."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("got to here cases 2")
	if queryInitiator.RoleID != "3" {
		resp.ResponseDescription = "Only a valid CSO is allowed to Initiate complaint request."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if NewCase.ClientType == "INDIVIDUAL" || NewCase.ClientType == "INDV" {
		NewCase.ClientType = "INDV"
	}
	if NewCase.ClientType == "CORPORATE" || NewCase.ClientType == "CORP" {
		NewCase.ClientType = "CORP"
	}
	refnum := utilities.GenFerenceNo()
	if len(refnum) < 16 {
		refnum = utilities.PadLeft(refnum, "0", 16)
	}
	fmt.Println("got to here cases 3")
	tblcase.RefNo = refnum
	tblcase.ClientType = NewCase.ClientType
	tblcase.FullName = NewCase.CFirstName + " " + NewCase.CLastName
	tblcase.CFirstName = NewCase.CFirstName
	tblcase.CMiddleName = NewCase.CMiddleName
	tblcase.CLastName = NewCase.CLastName
	tblcase.Address1 = NewCase.Address1
	if len(NewCase.Address2) > 50 {
		tblcase.Address2 = NewCase.Address2[0:50]
	} else if NewCase.Address2 == "" {
		tblcase.Address2 = "NA"
	} else {
		tblcase.Address2 = NewCase.Address2
	}

	tblcase.Country = "NG"

	tblcase.State = NewCase.State
	tblcase.City = NewCase.City
	tblcase.PosterCode = NewCase.PosterCode
	tblcase.CellPhone = NewCase.CellPhone
	tblcase.OfficePhone = NewCase.OfficePhone
	tblcase.Email = NewCase.Email
	tblcase.BranchName = NewCase.BranchName
	tblcase.AccountNumber = NewCase.AccountNumber
	tblcase.AccountType = NewCase.AccountType
	tblcase.AccountCurrency = NewCase.AccountCurrency
	tblcase.NameOfConsultant = NewCase.NameOfConsultant
	tblcase.ComplaintChannelID = NewCase.ComplaintChannelID //confirn  complaint channel id exist
	tblcase.Implication = NewCase.Implication
	tblcase.ComplaintLocation = NewCase.ComplaintLocation
	tblcase.Category = NewCase.Category //need to confirm category is not OTHERS
	tblcase.SubCategory = NewCase.SubCategory
	tblcase.Subject = NewCase.Subject
	tblcase.Description = NewCase.Description
	tblcase.AddedBy = queryInitiator.UserName
	tblcase.DisputeAmount = NewCase.DisputeAmount
	tblcase.AmountRefunded = NewCase.AmountRefunded
	tblcase.AmountRecovered = NewCase.AmountRecovered
	tblcase.ActionTaken = NewCase.ActionTaken
	tblcase.Prayer = NewCase.Prayer
	tblcase.Status = NewCase.Status
	tblcase.ComplaintType = "ATM"
	fmt.Println("got to here cases ")
	dt := time.Now()
	fmt.Println(dt.Format("DD-MM-YYYY"))
	tblcase.DateReceived = time.Now() //strconv.Itoa(year) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(day)
	if NewCase.Implication == "NON-FINANCIAL" {
		tblcase.Implication = "Y"
	} else {
		tblcase.Implication = "N"
	}
	tblcase.PreffeContPhone = NewCase.PreffeContPhone
	tblcase.PreffeEmail = NewCase.PreffeEmail
	tblcase.PreffeContact = NewCase.PreffeContact

	transactionDate, err := time.Parse("2006-01-02", NewCase.TransactionDate)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("got to here cases 22 ")
	tblcase.TransactionDate = transactionDate //time.Parse(NewCase.TransactionDate, time.Date())  //. NewCase.TransactionDate
	fmt.Println("got to here cases ")
	//log complaint
	logComplaint, err := compService.LogComplaint(ctx, tblcase)
	if logComplaint <= 0 {
		resp.ResponseCode = "01"
		resp.ResponseDescription = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.ResponseCode = "00"
	resp.ResponseDescription = "Compliant with reference ID " + tblcase.RefNo + " logged successfully."
	ctx.JSON(http.StatusOK, resp)
}

// GetComplaintByRefID godoc
// @Summary get logged complaint with reference ID
// @Produce json
// @Success 200 {object} models.TblCases
// @Router /complaint/GetComplaintByRefID/{ReferenceID} [get]
func GetComplaintByRefID(ctx *gin.Context) {
	fmt.Println("got to here cases 1")
	CompRefNo := ctx.Param("ReferenceID")
	if CompRefNo == "nil" {
		resp.ResponseDescription = "Reference number is required for this request."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("got to here cases 21")
	//get complaint with id
	CustCompliant, err := compService.GetSingleComplaintByComplaintID(ctx, CompRefNo)
	if err != nil {
		resp.ResponseDescription = "unable to fetch complaint at the moment. Please try again later."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusOK, resp)
		return
	}
	ctx.JSON(http.StatusOK, CustCompliant)
}

// CreateComplaintCategories godoc
// @Summary create new compliant
// @Produce json
// @Param user body models.ComplaintCategories true "Create new complaint category"
// @Success 200 {object} models.ResponseMessage
// @Router /complaint/CreateComplaintCategories/ [post]
func CreateComplaintCategories(ctx *gin.Context) {
	var req md.ComplaintCategories
	var reqBody md.TblComplaintCategories
	requestBody := ctx.ShouldBindJSON(&req)
	if requestBody != nil {
		ctx.JSON(http.StatusOK, "unable to read request at the moment. Please confirm the request body and try again")
		return
	}
	fmt.Println(req.CBNCategoryCode)
	//check it has not been created before
	checkCompl := compService.CheckCompliantCategory(ctx, req)
	if checkCompl > 0 {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "Complaint category already exist. Please re-check the name and try again."
		ctx.JSON(http.StatusOK, resp)
		return
	}
	reqBody.Category = req.Category
	reqBody.CBNCategoryCode = req.CBNCategoryCode
	//create category
	doCreate, err := compService.CreateCompliantCategory(ctx, reqBody)
	if err != nil {
		resp.ResponseDescription = err.Error()
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusOK, resp)
		return
	}
	if doCreate <= 0 {
		resp.ResponseDescription = "Service is unable to create category at the moment. Please try again later"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.ResponseDescription = "Complaint category created successfully."
	resp.ResponseCode = "00"
	ctx.JSON(http.StatusOK, resp)
}

// FetchComplaintCategories godoc
// @Summary fetch the list of complaint category
// @Produce json
// @Success 200 {object} models.TblComplaintCategories
// @Router /complaint/FetchComplaintCategories/ [get]
func FetchComplaintCategories(ctx *gin.Context) {
	CustCompliant := compService.FetchCompliantCategories(ctx)
	ctx.JSON(http.StatusOK, CustCompliant)
}
