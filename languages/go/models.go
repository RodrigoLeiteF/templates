package main

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
)

// UserModel is the model definition of a {{name}} document.
type {{name}Model struct {
	ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
	Status      string     `json:"status,omitempty" bson:"status,omitempty"`
	Avatar      string     `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Name        string     `json:"name,omitempty" bson:"name,omitempty"`
	Email       string     `json:"email,omitempty" bson:"email,omitempty"`
	Password    string     `json:"password,omitempty" bson:"password,omitempty"`
	Phone       string     `json:"phone,omitempty" bson:"phone,omitempty"`
	Role        string     `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	ActivatedAt *time.Time `json:"activated_at,omitempty" bson:"activated_at,omitempty"`
}
