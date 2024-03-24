package user

type User struct {
	name string // private attribute of the struct
	Age  int    // public attribute of the struct
}

// importers can call using user.New("")
func New(name string) *User {
	return &User{
		name: name,
	}
}

// importers can call using user.Name()
type UserInterface interface {
	// Name() returns the name of the user
	Name() string
}

func (u *User) Name() string {
	return u.name
}
