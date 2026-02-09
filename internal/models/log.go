package models

import "time"

type APILog struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	APIID          string    `bson:"api_id" json:"api_id"`
	Timestamp      time.Time `bson:"timestamp" json:"timestamp"`
	ResponseTimeMS int64     `bson:"response_time_ms" json:"response_time_ms"`
	StatusCode     int       `bson:"status_code" json:"status_code"`
	Success        bool      `bson:"success" json:"success"`
}
