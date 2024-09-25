package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Material struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Tags      []string           `bson:"tags" json:"tags"`
	Notes     []Note             `bson:"notes" json:"notes"`
	PDFs      []PDF              `bson:"pdfs" json:"pdfs"`
	Goals     []Goal             `bson:"goals" json:"goals"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type Note struct {
	NoteID    primitive.ObjectID `bson:"noteId,omitempty" json:"noteId,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type PDF struct {
	PDFID      primitive.ObjectID `bson:"pdfId,omitempty" json:"pdfId,omitempty"`
	Filename   string             `bson:"filename" json:"filename"`
	URL        string             `bson:"url" json:"url"`
	UploadedAt time.Time          `bson:"uploadedAt" json:"uploadedAt"`
}

type Goal struct {
	GoalID      primitive.ObjectID `bson:"goalId,omitempty" json:"goalId,omitempty"`
	Description string             `bson:"description" json:"description"`
	StartDate   time.Time          `bson:"startDate" json:"startDate"`
	EndDate     time.Time          `bson:"endDate" json:"endDate"`
	Status      string             `bson:"status" json:"status"`
}
