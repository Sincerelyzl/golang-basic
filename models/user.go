package models

type User struct {
	Name, Address string
}

func (u *User) UpdateAddress(address string) {
	u.Address = address
}
