package main

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
)

func validateGet{{name}}Request(r *{{name}}pb.GetUsersRequest) error {}
