package main

import (
	"time"

	"github.com/r3labs/diff"
)

type Entry struct {
	Id        string         `json:"id" bson:"_id"`
	EntityId  string         `json:"entity_id"`
	OldValue  string         `json:"old_value"`
	NewValue  string         `json:"new_value"`
	Changes   diff.Changelog `json:"changes"`
	UserId    string         `json:"user"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
}
