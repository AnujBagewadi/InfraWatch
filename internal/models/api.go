package models

import "time"

type API struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Name        string    `bson:"name" json:"name"`
	URL         string    `bson:"url" json:"url"`
	ThresholdMS int       `bson:"threshold_ms" json:"threshold_ms"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}
