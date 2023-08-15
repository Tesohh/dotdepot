package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Username          string             `bson:"username,omitempty"`
	PasswordEncrypted string             `bson:"password,omitempty"`
}
