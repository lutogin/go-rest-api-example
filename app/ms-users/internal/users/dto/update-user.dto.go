package usersDto

type UpdateUserDto struct {
	Id      string `json:"id" bson:"_id,omitempty" validate:"required,mongodb"`
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email" validate:"email"`
	Phone   string `json:"phone" bson:"phone" validate:"e164"`
	PwdHash string `json:"-" bson:"pwdHash"`
}
