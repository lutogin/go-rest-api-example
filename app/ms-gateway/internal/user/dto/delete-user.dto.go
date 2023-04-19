package userDto

type DeleteUserDto struct {
	Id string `json:"id" bson:"_id,omitempty"` // omitempty means can be null
}
