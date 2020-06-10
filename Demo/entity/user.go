package entity

type User struct {
	Firstname	string `json:"firstname,omitempty"`
	Lastname	string `json:"lastname,omitempty"`
	Username	string `json:"username,omitempty"`
	Password	string `json:"password"`
	PhoneNumber	string `json:"phoneNumber,omitempty"`
	Email		string `json:"email,omitempty"`
	IsAdmin		bool   `json:"isAdmin"`
}
