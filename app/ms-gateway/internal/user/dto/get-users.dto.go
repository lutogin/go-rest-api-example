package userDto

type GetUsersDto struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
}
