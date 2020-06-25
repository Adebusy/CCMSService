package models

//Roles all roles
type Roles struct {
	Roles []Role `json:"roles"`
}

// Role struct which contains a name and ID
type Role struct {
	Name string `json:"name"`
	ID   int    `json:"ID"`
}
