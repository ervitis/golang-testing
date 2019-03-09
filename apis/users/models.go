package users

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Surname *string `json:"surname"`
	Email   string  `json:"email"`
	Gender  *string `json:"gender"`
	Country string  `json:"country"`
}
