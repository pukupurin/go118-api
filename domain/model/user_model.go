package model

type User struct {
	ID   int
	Name string
	Age  int
}
type Users []User

// NewUser User のコンストラクタ
func NewUser(name string, age int) (*User, error) {

	user := &User{
		Name: name,
		Age:  age,
	}

	return user, nil
}
