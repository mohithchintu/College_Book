package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"passwordHash" json:"passwordHash"`
	Email       string             `bson:"email" json:"email"`
	Profile     Profile            `bson:"profile" json:"profile"`
	Preferences Preferences        `bson:"preferences" json:"preferences"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Profile struct {
	FirstName      string    `bson:"firstName" json:"firstName"`
	LastName       string    `bson:"lastName" json:"lastName"`
	DateOfBirth    time.Time `bson:"dateOfBirth" json:"dateOfBirth"`
	ProfilePicture string    `bson:"profilePicture" json:"profilePicture"`
}

type Preferences struct {
	NotificationSettings NotificationSettings `bson:"notificationSettings" json:"notificationSettings"`
	Theme                string               `bson:"theme" json:"theme"`
}

type NotificationSettings struct {
	Email bool `bson:"email" json:"email"`
	App   bool `bson:"app" json:"app"`
}
