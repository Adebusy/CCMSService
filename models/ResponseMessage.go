package models

//ResponseMessage general response
type ResponseMessage struct {
	ResponseDescription string `json:"ResponseDescription,omitempty"`
	ResponseCode        string `json:"ResponseCode,omitempty"`
}
