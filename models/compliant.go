package models

import (
	"time"
)

//TblCases complaint schema
type TblCases struct {
	ID                 int       `json:"ID" validate:"omitempty"`
	RefNo              string    `json:"RefNo" validate:"omitempty"`
	ClientType         string    `json:"ClientType" validate:"omitempty"`
	FullName           string    `json:"FullName" validate:"omitempty"`
	CFirstName         string    `json:"CFirstName" validate:"omitempty"`
	CMiddleName        string    `json:"CMiddleName" validate:"omitempty"`
	CLastName          string    `json:"CLastName" validate:"omitempty"`
	Address1           string    `json:"Address1" validate:"omitempty"`
	Address2           string    `json:"Address2" validate:"omitempty"`
	Country            string    `json:"Country" validate:"omitempty"`
	State              string    `json:"State" validate:"omitempty"`
	City               string    `json:"City" validate:"omitempty"`
	PosterCode         string    `json:"PosterCode" validate:"omitempty"`
	CellPhone          string    `json:"CellPhone" validate:"omitempty"`
	OfficePhone        string    `json:"OfficePhone" validate:"omitempty"`
	Email              string    `json:"Email" validate:"omitempty"`
	BranchName         string    `json:"BranchName" validate:"omitempty"`
	AccountNumber      string    `json:"AccountNumber" validate:"omitempty"`
	AccountType        string    `json:"AccountType" validate:"omitempty"`
	AccountCurrency    string    `json:"AccountCurrency" validate:"omitempty"`
	NameOfConsultant   string    `json:"NameOfConsultant" validate:"omitempty"`
	ComplaintChannelID string    `json:"ComplaintChannelID" validate:"omitempty"`
	Implication        string    `json:"Implication" validate:"omitempty"`
	ComplaintLocation  string    `json:"ComplaintLocation" validate:"omitempty"`
	Category           string    `json:"Category" validate:"omitempty"`
	SubCategory        string    `json:"SubCategory" validate:"omitempty"`
	Subject            string    `json:"Subject" validate:"omitempty"`
	Description        string    `json:"Description" validate:"omitempty"`
	Prayer             string    `json:"Prayer" validate:"omitempty"`
	DateReceived       time.Time `json:"DateReceived" validate:"omitempty"`
	DisputeAmount      string    `json:"DisputeAmount" validate:"omitempty"`
	AmountRefunded     string    `json:"AmountRefunded" validate:"omitempty"`
	AmountRecovered    string    `json:"AmountRecovered" validate:"omitempty"`
	Remarks            string    `json:"Remarks" validate:"omitempty"`
	DateClosed         time.Time `json:"DateClosed" validate:"omitempty"`
	Status             string    `json:"Status" validate:"omitempty"`
	ActionTaken        string    `json:"ActionTaken" validate:"omitempty"`
	AddedBy            string    `json:"AddedBy" validate:"omitempty"`
	updatedby          string    `json:"updatedby" validate:"omitempty"`
	ComplaintType      string    `json:"ComplaintType" validate:"omitempty"`
	comment            string    `json:"comment" validate:"omitempty"`
	PreffeContPhone    string    `json:"PreffeContPhone" validate:"omitempty"`
	PreffeEmail        string    `json:"PreffeEmail" validate:"omitempty"`
	PreffeContact      string    `json:"PreffeContact" validate:"omitempty"`
	TransactionDate    time.Time `json:"TransactionDate" validate:"omitempty"`
	DateCreated        time.Time `json:"DateCreated" validate:"omitempty"`
}

//Cases obj
type Cases struct {
	RefNo              string    `json:"RefNo" validate:"omitempty"`
	ClientType         string    `json:"ClientType" validate:"omitempty"`
	FullName           string    `json:"FullName" validate:"omitempty"`
	CFirstName         string    `json:"CFirstName" validate:"omitempty"`
	CMiddleName        string    `json:"CMiddleName" validate:"omitempty"`
	CLastName          string    `json:"CLastName" validate:"omitempty"`
	Address1           string    `json:"Address1" validate:"omitempty"`
	Address2           string    `json:"Address2" validate:"omitempty"`
	Country            string    `json:"Country" validate:"omitempty"`
	State              string    `json:"State" validate:"omitempty"`
	City               string    `json:"City" validate:"omitempty"`
	PosterCode         string    `json:"PosterCode" validate:"omitempty"`
	CellPhone          string    `json:"CellPhone" validate:"omitempty"`
	OfficePhone        string    `json:"OfficePhone" validate:"omitempty"`
	Email              string    `json:"Email" validate:"omitempty"`
	BranchName         string    `json:"BranchName" validate:"omitempty"`
	AccountNumber      string    `json:"AccountNumber" validate:"omitempty"`
	AccountType        string    `json:"AccountType" validate:"omitempty"`
	AccountCurrency    string    `json:"AccountCurrency" validate:"omitempty"`
	NameOfConsultant   string    `json:"NameOfConsultant" validate:"omitempty"`
	ComplaintChannelID string    `json:"ComplaintChannelID" validate:"omitempty"`
	Implication        string    `json:"Implication" validate:"omitempty"`
	ComplaintLocation  string    `json:"ComplaintLocation" validate:"omitempty"`
	Category           string    `json:"Category" validate:"omitempty"`
	SubCategory        string    `json:"SubCategory" validate:"omitempty"`
	Subject            string    `json:"Subject" validate:"omitempty"`
	Description        string    `json:"Description" validate:"omitempty"`
	Prayer             string    `json:"Prayer" validate:"omitempty"`
	DateReceived       time.Time `json:"DateReceived" validate:"omitempty"`
	DisputeAmount      string    `json:"DisputeAmount" validate:"omitempty"`   //decimal.Decimal
	AmountRefunded     string    `json:"AmountRefunded" validate:"omitempty"`  //decimal.Decimal
	AmountRecovered    string    `json:"AmountRecovered" validate:"omitempty"` //decimal.Decimal
	Remarks            string    `json:"Remarks" validate:"omitempty"`
	Status             string    `json:"Status" validate:"omitempty"`
	ActionTaken        string    `json:"ActionTaken" validate:"omitempty"`
	AddedBy            string    `json:"AddedBy" validate:"omitempty"`
	ComplaintType      string    `json:"ComplaintType" validate:"omitempty"`
	comment            string    `json:"comment" validate:"omitempty"`
	PreffeContPhone    string    `json:"PreffeContPhone" validate:"omitempty"`
	PreffeEmail        string    `json:"PreffeEmail" validate:"omitempty"`
	PreffeContact      string    `json:"PreffeContact" validate:"omitempty"`
	TransactionDate    time.Time `json:"TransactionDate" validate:"omitempty"`
}

type ComplaintStatus struct {
	StatusName string `json:"StatusName" validate:"omitempty"`
	StatusID   string `json:"StatusID" validate:"omitempty"`
}

type TblComplaintSubCategories struct {
	ID          int    `json:"ID" validate:"omitempty"`
	SubCategory string `json:"SubCategory" validate:"omitempty"`
}

type TblComplaintChannels struct {
	ID        int       `json:"ID" validate:"omitempty"`
	Name      string    `json:"Name" validate:"omitempty"`
	DateAdded time.Time `json:"DateAdded" validate:"omitempty"`
	Audit     string    `json:"Audit" validate:"omitempty"`
}
