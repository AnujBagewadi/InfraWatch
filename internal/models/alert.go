package models

import "time"

type Alert struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	APIID     string    `bson:"api_id" json:"api_id"`
	Message   string    `bson:"message" json:"message"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Resolved  bool      `bson:"resolved" json:"resolved"`
}
