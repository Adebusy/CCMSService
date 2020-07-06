package models

type TblUsers struct {
	ID             int    `json:"ID" validate:"omitempty"`
	UserName       string `json:"UserName" validate:"omitempty"`
	UserStatus     string `json:"UserStatus" validate:"omitempty"`
	Email          string `json:"Email" validate:"omitempty"`
	FirstName      string `json:"FirstName" validate:"omitempty"`
	LastName       string `json:"LastName" validate:"omitempty"`
	DateAdded      string `json:"DateAdded" validate:"omitempty"`
	RoleID         string `json:"RoleID" validate:"omitempty"`
	UpdatedBy      string `json:"UpdatedBy" validate:"omitempty"`
	OfficeLocation string `json:"OfficeLocation" validate:"omitempty"`
}

type TblRoles struct {
	ID       int    `json:"ID" validate:"omitempty"`
	UserType string `json:"UserType" validate:"omitempty"` // "CUSTOMER SERVICE OFFICER", hop, Admin
}

type User struct {
	UserName       string `json:"UserName" validate:"omitempty"`
	UserStatus     string `json:"UserStatus" validate:"omitempty"`
	Email          string `json:"Email" validate:"omitempty"`
	FirstName      string `json:"FirstName" validate:"omitempty"`
	MiddleName     string `json:"MiddleName" validate:"omitempty"`
	LastName       string `json:"LastName" validate:"omitempty"`
	DateAdded      string `json:"DateAdded" validate:"omitempty"`
	RoleID         string `json:"RoleID" validate:"omitempty"`
	UpdatedBy      string `json:"UpdatedBy" validate:"omitempty"`
	OfficeLocation string `json:"OfficeLocation" validate:"omitempty"`
}
