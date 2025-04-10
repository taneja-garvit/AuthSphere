package models

type User struct{
	ID       string `json:"id" bson:"_id,omitempty"` // omitempty is used to make the field optional
	Email string `json:"email" bson:"email" validate:"required,email"` // validate is used to
	Password string `json:"password" bson:"password"`
}

// json  used for transmitting data between a server and a client (such as in REST APIs).
// bson is used for storing data in MongoDB and exchanging data. It is similar to JSON but is more efficient for storage and speed because it uses binary encoding