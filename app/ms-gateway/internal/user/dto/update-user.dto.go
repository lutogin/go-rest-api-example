package userDto

type UpdateUserDto struct {
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Phone   string `json:"phone" bson:"phone"`
	PwdHash string `json:"-" bson:"pwdHash"`
}
