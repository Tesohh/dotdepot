package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Paths struct {
	Windows string `bson:"windows,omitempty"`
	MacOS   string `bson:"macos,omitempty"`
	Linux   string `bson:"linux,omitempty"`
}

type Dotfile struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FileName string             `bson:"filename,omitempty"`
	UserName string             `bson:"username,omitempty"`
	Content  string             `bson:"content,omitempty"`
	Paths    Paths              `bson:"paths,omitempty"`
}
