package controllers

type user struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Surname *string `json:"surname"`
	Email   string  `json:"email"`
	Gender  *string `json:"gender"`
	Country string  `json:"country"`
}

type company struct {
	Id          int    `json:"id"`
	CompanyName string `json:"companyName"`
	UserId      int    `json:"userId"`
	Country     string `json:"country"`
	City        string `json:"city"`
}
