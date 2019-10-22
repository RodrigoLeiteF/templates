package main

import (
	validation "github.com/go-ozzo/ozzo-validation"
	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
	v "github.com/upnid/validations/go"
)

func ValidateCreateEntry(req {{name}}pb.CreateEntryRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.EntityId, validation.Required),
		validation.Field(&req.EntityId, validation.By(v.IsAnyID)),
		validation.Field(&req.UserId, validation.Required),
		validation.Field(&req.OldValue, validation.Required),
		validation.Field(&req.NewValue, validation.Required),
	)
}

func ValidateGetEntity(req {{name}}pb.GetEntityRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.EntityId, validation.Required),
	)
}
