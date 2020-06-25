package utilities

import (
	"fmt"
	"math/rand"
	"regexp"

	"github.com/Adebusy/CCMSService/features/users"
	modules "github.com/Adebusy/CCMSService/models"
)

//ValidateUserRequest validates user request
func ValidateUserRequest(requestBody users.User) modules.ResponseMessage {
	var resp modules.RequestResponse
	resp.ResponseCode = ""
	resp.ResponseMessage = ""
	if !ValidateEmail(requestBody.Email) {
		resp.ResponseMessage = "Email address must be valid"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.FirstName == "" {
		resp.ResponseMessage = "FirstName is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.LastName == "" {
		resp.ResponseMessage = "Lastname is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.MiddleName == "" {
		resp.ResponseMessage = "Middlename is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.OfficeLocation == "" {
		resp.ResponseMessage = "Office location is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.RoleID == "" {
		resp.ResponseMessage = "Role ID is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.UserName == "" {
		resp.ResponseMessage = "Username is required"
		resp.ResponseCode = "01"
		return resp
	}

	return resp
}

//ValidateEmail used to validate email address
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

func init() {
	//connect = con.NewClient()
	DbGorm, ErrGorm = sqlserver.ConnectGorm()
	if ErrGorm != nil {
		fmt.Printf(ErrGorm.Error())
	}
}

//GenerateOTP generates random number
func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numbers[rand.Int63()%int64(len(numbers))]
	}
	return string(b)
}
