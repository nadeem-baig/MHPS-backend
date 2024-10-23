package types

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=20"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	CreatedAt string `json:"created_at"`
}


type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id string) (*User, error)
	CreateUser(User) error
}

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	CreatedAt string `json:"created_at"`
}

type ProductStore interface {
	GetProducts() ([]Product, error)
}


//Members Registration form

type Member struct {
	AadhaarNumber string    `json:"aadhaarNumber"`
	Address       string    `json:"address"`
	BloodGroup    string    `json:"bloodGroup"`
	ContactNumber string    `json:"contactNumber"`
	DateOfBirth   string `json:"dateOfBirth"`
	Education     string    `json:"education"`
	Email         string    `json:"email"`
	FatherName    string    `json:"fatherName"`
	MaritalStatus string    `json:"maritalStatus"`
	Name          string    `json:"name"`
	StdPin        string    `json:"stdPin"`
}

type RegisterMemberPayload = Member