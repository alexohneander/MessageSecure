package models

import "time"

type Message struct {
	Id        string    `bson:"id"`
	Message   string    `bson:"message"`
	CreatedAt time.Time `bson:"created_at"`
}
