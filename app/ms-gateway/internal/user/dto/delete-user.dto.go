package userDto

type DeleteUserDto struct {
	ID string `json:"id" bson:"_id,omitempty"` // omitempty means can be null
}
