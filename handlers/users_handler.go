package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Adebusy/CCMSService/models"
	md "github.com/Adebusy/CCMSService/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"

	ut "github.com/Adebusy/CCMSService/utilities"
)

//"github.com/Adebusy/CCMSService/datastore/sqlserver"
//"github.com/Adebusy/CCMSService/driver"

var resp models.ResponseMessage

// FetchRoles godoc
// @Summary create new question
// @Produce json
// @Success 200 {object} models.Roles
// @Router /Users/FetchAvailableRoles/ [get]
func FetchRoles(ctx *gin.Context) {
	resp.ResponseDescription = ""
	resp.ResponseCode = ""
	ctx.JSON(http.StatusOK, GetAvailableRoles().Roles)
}

// CreateUser godoc
// @Summary create new question
// @Produce json
// @Param user body models.User true "create new user"
// @Success 200 {object} models.ResponseMessage
// @Router /Users/CreateUser/ [post]
func CreateUser(ctx *gin.Context) {
	var RequestBody md.User
	resp.ResponseDescription = ""
	resp.ResponseCode = ""
	if userObj := ctx.ShouldBindJSON(&RequestBody); userObj != nil {
		httputil.NewError(ctx, http.StatusBadRequest, userObj)
		ctx.JSON(http.StatusBadRequest, userObj)
		return
	}

	//validate user request
	valRequest := ValidateUserRequest(RequestBody)
	if valRequest.ResponseCode != "" {
		ctx.JSON(http.StatusBadRequest, valRequest)
		return
	}
	//validate user request
	usr, err := UserInterface.CheckUser(ctx, RequestBody.UserName)
	if err != nil {
		resp.ResponseDescription = "Unable to validate user at the moment. Please try again later!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	if usr.ID > 0 {
		resp.ResponseDescription = "Username already exist. Please re-confirm and try again!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//validate role ID
	AllRole := GetAvailableRoles()
	var isroleValid bool = false

	for i := 0; i < len(AllRole.Roles); i++ {
		if strconv.Itoa(AllRole.Roles[i].ID) == RequestBody.RoleID {
			isroleValid = true
			break
		}
	}

	if !isroleValid {
		resp.ResponseDescription = "Role ID must be within available role ID in the FetchAvailableRoles method."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//create user
	var tbluserObj = md.TblUsers{DateAdded: time.Now().String(), UserName: RequestBody.UserName, UserStatus: "Active", Email: RequestBody.Email, FirstName: RequestBody.FirstName, LastName: RequestBody.LastName, RoleID: RequestBody.RoleID, UpdatedBy: "new user", OfficeLocation: RequestBody.OfficeLocation}
	fmt.Println(tbluserObj.LastName)
	createusr := UserInterface.CreateUser(ctx, tbluserObj)
	ctx.JSON(http.StatusOK, createusr)
}

// ChangeUserDetails godoc
// @Summary create new question
// @Produce json
// @Param user body models.User true "create new user"
// @Success 200 {object} models.ResponseMessage
// @Router /Users/ChangeUserDetails/ [post]
func ChangeUserDetails(ctx *gin.Context) {
	var resp md.ResponseMessage
	var RequestBody models.User
	resp.ResponseDescription = ""
	resp.ResponseCode = ""
	ctx.ShouldBindJSON(&RequestBody)

	//validate user request
	valRequest := ValidateUserRequest(RequestBody)
	if valRequest.ResponseCode != "" {
		ctx.JSON(http.StatusBadRequest, valRequest)
		return
	}

	//validate user request
	usr, err := UserInterface.CheckUser(ctx, RequestBody.UserName)
	if err != nil {
		resp.ResponseDescription = "Unable to validate user at the moment. Please try again later!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if usr.ID <= 0 {
		resp.ResponseDescription = "Email address does not exist. Please re-confirm and try again!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	//do update
	var tbluserObj = md.TblUsers{ID: usr.ID, DateAdded: "timer", UserName: usr.UserName, UserStatus: "Active", Email: usr.Email, FirstName: RequestBody.FirstName, LastName: RequestBody.LastName, RoleID: "1", UpdatedBy: "new user", OfficeLocation: usr.OfficeLocation}
	fmt.Println(tbluserObj.LastName)
	fmt.Println("last name")
	updateuserReq := UserInterface.UpdateUser(ctx, tbluserObj)
	ctx.JSON(http.StatusOK, updateuserReq)
}

// ChangeUserStatus godoc
// @Summary Update user status ('Active', 'Non-Active')
// @Produce json
// @Success 200 {object} models.ResponseMessage
// @Router /Users/ChangeUserStatus/{Username}/{Status}/{RoleID} [get]
func ChangeUserStatus(ctx *gin.Context) {

	requester := ctx.GetHeader("RequestByUsername")
	if requester == "" {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "RequestByUsername header is required."
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	//check request id

	usrRequest, errusr := UserInterface.CheckUser(ctx, requester)
	if errusr != nil {
		resp.ResponseDescription = "Unable to validate requester re-check the header. Please try again later!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	if usrRequest.ID <= 0 {
		resp.ResponseDescription = "Initiator does not exist. Please try again later!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	if usrRequest.RoleID != "1" && usrRequest.RoleID != "2" {
		resp.ResponseDescription = "Initiator must be either Admin or super-Admin to perform this function" + usrRequest.RoleID
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if Username := ctx.Param("Username"); Username == "" {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "username is required in the request parameter for this operation."
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if NewStatus := ctx.Param("Status"); NewStatus == "" {
		resp.ResponseCode = "01"
		resp.ResponseDescription = "Proposed Status ('Active', 'Non-Active') is required for this operation."
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	// //confirm user exists
	// //validate user request
	usr, err := UserInterface.CheckUser(ctx, ctx.Param("Username"))
	if err != nil {
		resp.ResponseDescription = "Unable to validate user at the moment. Please try again later!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if usr.ID <= 0 {
		resp.ResponseDescription = "user does not exist. Please re-confirm and try again!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	usr.UserStatus = ctx.Param("Status")
	usr.RoleID = ctx.Param("RoleID")
	usr.UpdatedBy = requester
	statusReq := UserInterface.ChangeUserStatusRole(ctx, usr)
	ctx.JSON(http.StatusOK, statusReq)
}

// GetUserFullInfoByEmail godoc
// @Summary get user detials by user email address
// @Produce json
// @Param UserEmail path string true "User email address"
// @Success 200 {object} models.User
// @Router /Users/GetUserFullInfoByEmail/{UserEmail} [get]
func GetUserFullInfoByEmail(ctx *gin.Context) {
	fmt.Println("print")
	userEmail := ctx.Param(`UserEmail`)
	var resp md.ResponseMessage
	//check user email address
	if ut.ValidateEmail(userEmail) == false {
		resp.ResponseDescription = "Invalid Email address supplied. Please re-confirm."
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("print1")
	queryResp, er := UserInterface.GetUserByEmail(ctx, userEmail)
	if er != nil {
		resp.ResponseDescription = "Error occurred. Please try again later!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("print2")
	if queryResp.ID <= 0 {
		resp.ResponseDescription = "User does not exist. Please re-confirm the email address!!!"
		resp.ResponseCode = "01"
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	fmt.Println("print3")
	var respBody md.User
	respBody.UpdatedBy = queryResp.UpdatedBy
	respBody.Email = queryResp.Email
	respBody.FirstName = queryResp.FirstName
	respBody.LastName = queryResp.LastName
	respBody.MiddleName = " "
	respBody.OfficeLocation = queryResp.OfficeLocation
	respBody.RoleID = queryResp.RoleID
	respBody.UserName = queryResp.UserName
	ctx.JSON(http.StatusBadRequest, respBody)
	return
}

//GetAvailableRoles all roles
func GetAvailableRoles() md.Roles {
	jsonFile, err := os.Open("roles.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	var roles md.Roles
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &roles)
	return roles
}
