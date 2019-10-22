package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/lucsky/cuid"
	"github.com/r3labs/diff"
	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	clients Clients
}

type Clients struct {
	Db *mongo.Database
}

const collection string = "{{name}}"

func (s *Server) CreateEntry(ctx context.Context, request *{{name}}pb.CreateEntryRequest) (*{{name}}pb.CreateEntryResponse, error) {
	if err := ValidateCreateEntry(*request); err != nil {
		return nil, status.Error(codes.InvalidArgument, "VALIDATION_ERROR: "+err.Error())
	}

	entity_id := request.EntityId
	user_id := request.UserId
	old_value := request.OldValue
	new_value := request.NewValue

	if err1 != nil || err2 != nil {
		return nil, errors.New("Serialization error")
	}

	diffs, err := CompareObjects(old_value, new_value)
	if err != nil {
		return nil, err
	}

	entry := &Entry{
		Id:        "en_" + cuid.New(),
		EntityId:  entity_id,
		OldValue:  old_value_str,
		NewValue:  new_value_str,
		Changes:   diffs,
		UserId:    user_id,
		CreatedAt: time.Now(),
	}

	_, err = s.clients.Db.Collection(collection).InsertOne(ctx, entry)
	if err != nil {
		return nil, status.New(codes.Internal, "LOG_MONGO_INSERT_ERROR").Err()
	}

	createdat, _ := ptypes.TimestampProto(entry.CreatedAt)
	var changes []*{{name}}pb.Change
	for _, change := range entry.Changes {
		changes = append(changes, ChangeToMessage(change))
	}

	returned_entry := &{{name}}pb.Entry{
		Id:         entry.Id,
		EntityId:   entry.EntityId,
		OldValue:   entry.OldValue,
		NewValue:   entry.NewValue,
		Changes:    changes,
		UserId:     entry.UserId,
		CreateTime: createdat,
	}

	return &{{name}}pb.CreateEntryResponse{
		Entry: returned_entry,
	}, nil
}
