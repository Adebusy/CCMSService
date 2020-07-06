package utilities

import (
	md "github.com/Adebusy/CCMSService/models"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var resp md.ResponseMessage

//ValidateUserRequest validates user request
func ValidateUserRequest(requestBody md.User) md.ResponseMessage {
	resp.ResponseCode = ""
	resp.ResponseDescription = ""
	if !ValidateEmail(requestBody.Email) {
		resp.ResponseDescription = "Email address must be valid"
		resp.ResponseCode = "01"
		return resp
	}
	//"github.com/Adebusy/CCMSService/features/users"
	if requestBody.FirstName == "" {
		resp.ResponseDescription = "FirstName is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.LastName == "" {
		resp.ResponseDescription = "Lastname is required"
		resp.ResponseCode = "01"
		return resp
	}

	// if requestBody.MiddleName == "" {
	// 	resp.ResponseMessage = "Middlename is required"
	// 	resp.ResponseCode = "01"
	// 	return resp
	// }

	if requestBody.OfficeLocation == "" {
		resp.ResponseDescription = "Office location is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.RoleID == "" {
		resp.ResponseDescription = "Role ID is required"
		resp.ResponseCode = "01"
		return resp
	}

	if requestBody.UserName == "" {
		resp.ResponseDescription = "Username is required"
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

// func init() {
// 	//connect = con.NewClient()
// 	DbGorm, ErrGorm = sqlserver.ConnectGorm()
// 	if ErrGorm != nil {
// 		fmt.Printf(ErrGorm.Error())
// 	}
// }
var numbers = "0123456789"

//GenerateOTP generates random number
func GenerateOTP(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numbers[rand.Int63()%int64(len(numbers))]
	}
	return string(b)
}

//GenFerenceNo generates ference number for complaint
func GenFerenceNo() string {
	year, month, day := time.Now().Date()
	sec := time.Now().Second
	refnum := strconv.Itoa(year) + strconv.Itoa(int(month)) + strconv.Itoa(day) + strconv.Itoa(sec())
	refnum = "SBN" + refnum + GenerateOTP(4)
	return refnum
}

//GoDotEnvVariable load env file
func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		//log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
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

//ValidateComplaintRequest call to confirm request parameters
func ValidateComplaintRequest(NewCase md.Cases) md.ResponseMessage {
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
	return resp
}
