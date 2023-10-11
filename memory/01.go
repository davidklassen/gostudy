package memory

type User struct {
	FirstName string
	LastName  string
	FullName  string
}

func SetFullName(u *User) {
	u.FullName = u.FirstName + " " + u.LastName
}
