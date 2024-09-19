package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Memories  []MemoryDetail     `bson:"memories" json:"memories"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type MemoryDetail struct {
	Type        string        `bson:"type" json:"type"`
	Content     string        `bson:"content" json:"content"`
	DateCreated time.Time     `bson:"dateCreated" json:"dateCreated"`
	Tags        []string      `bson:"tags" json:"tags"`
	Media       []MediaDetail `bson:"media" json:"media"`
}

type MediaDetail struct {
	MediaID    primitive.ObjectID `bson:"mediaId,omitempty" json:"mediaId,omitempty"`
	MediaType  string             `bson:"mediaType" json:"mediaType"`
	Filename   string             `bson:"filename" json:"filename"`
	URL        string             `bson:"url" json:"url"`
	UploadedAt time.Time          `bson:"uploadedAt" json:"uploadedAt"`
}
