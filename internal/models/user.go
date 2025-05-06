package models

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	JobTitle  *string `json:"jobTitle,omitempty"`
	CompanyID int     `json:"companyId"`
	Email     string  `json:"email"`
	Lemons    int     `json:"lemons"`
	Diamonds  int     `json:"diamonds"`
	UserRole  string  `json:"userRole"`
	IsActive  bool    `json:"isActive"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	JobTitle  *string `json:"jobTitle"`
	CompanyID *int    `json:"companyId"`
	Email     *string `json:"email"`
	Lemons    *int    `json:"lemons"`
	Diamonds  *int    `json:"diamonds"`
	UserRole  *string `json:"userRole"`
	IsActive  *bool   `json:"isActive"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
