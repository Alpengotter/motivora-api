package models

type EmployerStatistic struct {
	Users    int `json:"users"`
	Lemons   int `json:"lemons"`
	Diamonds int `json:"diamonds"`
}

type CompaniesStatistic struct {
	Companies int `json:"companies"`
	Lemons    int `json:"lemons"`
	Diamonds  int `json:"diamonds"`
}
