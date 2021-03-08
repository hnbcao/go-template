package entity

import (
	"context"
	"time"
)

type (
	// User represents a cluster of the system.
	User struct {
		ID         int64     `json:"user_id,omitempty"`
		Name       string    `json:"username,omitempty"`
		Password   string    `json:"password,omitempty"`
		Salt       string    `json:"salt,omitempty"`
		Status     bool      `json:"status,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
	}

	// UserStore defines operations for working with user.
	UserStore interface {

		// Find returns a cluster from the datastore.
		Find(context.Context, int64) (*User, error)

		FindWithName(context.Context, string) (*User, error)

		Create(context.Context, *User) error

		Update(context.Context, *User) error
	}
)
