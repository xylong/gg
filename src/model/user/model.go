package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" binding:"min=4"`
}

func NewUser(attrs ...UserAttrFunc) *User {
	user := &User{}
	UserAttrFuncs(attrs).Apply(user)
	return user
}
