package entities

type User struct {
	ID       string `json:"userid" gorm:"primarykey"`
	Name     string `json:"username"`
	Email    string `json:"useremail" gorm:"unique"`
	Password string `json:"-"`
}
