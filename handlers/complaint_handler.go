package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Adebusy/CCMSService/driver"
	"github.com/Adebusy/CCMSService/utilities"

	"github.com/Adebusy/CCMSService/models"
	"github.com/gin-gonic/gin"

	"github.com/Adebusy/CCMSService/datastore/sqlserver"
)

var compService = sqlserver.NewComplaintService(driver.GetDB())

// LogComplaintRequest godoc
// @Summary create new compliant
// @Produce json
// @Param user body models.Cases true "Create new complaint"
// @Success 200 {object} models.ResponseMessage
// @Router /complaint/LogComplaintRequest/ [post]
func LogComplaintRequest(ctx *gin.Context) {
	fmt.Println("got to here cases 1")
	var NewCase models.Cases
	var resp models.ResponseMessage
	var tblcase models.TblCases

	err := ctx.ShouldBindJSON(&NewCase)
	if err != nil {
		resp.ResponseDescription = err.Error()
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	if NewCase.AddedBy == "" {
		resp.ResponseDescription = "Initiators name is required for this request to be treated."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	//confirm initiators name
	queryInitiator, _ := UserInterface.CheckUser(ctx, NewCase.AddedBy)
	if queryInitiator.ID <= 0 {
		resp.ResponseDescription = "Initiators does not exist."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("got to here cases 2")
	if queryInitiator.RoleID != "CSO" {
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
	year, month, day := time.Now().Date()
	sec := time.Now().Second
	refnum := strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day) + strconv.Itoa(sec)
	refnum = "SBN" + refnum + utilities.GenerateOTP(4)

	if len(refnum) < 16 {
		refnum = PadLeft(refnum, "0", 16)
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
	tblcase.DateReceived = strconv.Itoa(year) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(day)
	if NewCase.Implication == "NON-FINANCIAL" {
		tblcase.Implication = "Y"
	} else {
		tblcase.Implication = "N"
	}
	tblcase.PreffeContPhone = NewCase.PreffeContPhone
	tblcase.PreffeEmail = NewCase.PreffeEmail
	tblcase.PreffeContact = NewCase.PreffeContact
	tblcase.TransactionDate = NewCase.TransactionDate
	fmt.Println("got to here cases ")
	//log complaint
	logComplaint, err := compService.LogComplaint(ctx, tblcase)
	if logComplaint > 0 {
		resp.ResponseCode = "01"
		resp.ResponseDescription = err.Error()
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp.ResponseCode = "01"
	resp.ResponseDescription = err.Error()
	ctx.JSON(http.StatusOK, resp)
	return
}

//PadLeft pad reference number
func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

var resp models.ResponseMessage

//validateComplaintRequest call to confirm request parameters
func validateComplaintRequest(NewCase models.Cases) models.ResponseMessage {
	//resp models.ResponseMessage
	if NewCase.Description == "" {
		resp.ResponseDescription = "Complaint description must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.ComplaintChannelID == "" {
		resp.ResponseDescription = "Complaint channel must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.Category == "" {
		resp.ResponseDescription = "Complaint Category must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.SubCategory == "" {
		resp.ResponseDescription = "Complaint sub category must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.Subject == "" {
		resp.ResponseDescription = "Subject of complaint must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.ClientType == "" {
		resp.ResponseDescription = "Subject of complaint must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.Category == "OTHERS" {
		resp.ResponseDescription = "complaint category must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.SubCategory == "" {
		resp.ResponseDescription = "complaint sub=category must be selected."
		resp.ResponseCode = "01"
		return resp
	}
	if NewCase.Subject == "" {
		resp.ResponseDescription = "Subject must be selected."
		resp.ResponseCode = "01"
		return resp
	}
}
