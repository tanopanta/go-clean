package domain

// User は ユーザーの構造体です
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Users は ユーザーの集まりです
type Users []User
