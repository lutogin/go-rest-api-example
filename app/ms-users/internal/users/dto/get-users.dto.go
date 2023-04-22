package usersDto

type GetUsersDto struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email" validate:"email"`
	Phone string `json:"phone" bson:"phone"`
}
