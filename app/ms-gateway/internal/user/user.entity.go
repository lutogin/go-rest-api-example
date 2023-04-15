package user

type UserEntity struct {
	ID      string `json:"id" bson:"_id,omitempty"` // omitempty means can be null
	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Phone   string `json:"phone" bson:"phone"`
	PwdHash string `json:"-" bson:"pwdHash"` // json:"-" means don't be get in JSON
}
