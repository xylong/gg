package user

type UserAttrFunc func(user *User)
type UserAttrFuncs []UserAttrFunc

func WithId(id int) UserAttrFunc {
	return func(user *User) {
		user.ID = id
	}
}

func WithName(name string) UserAttrFunc {
	return func(user *User) {
		user.Name = name
	}
}

func (u UserAttrFuncs) Apply(user *User) {
	for _, f := range u {
		f(user)
	}
}

func (u *User) Mutate(attrs ...UserAttrFunc) *User {
	UserAttrFuncs(attrs).Apply(u)
	return u
}
