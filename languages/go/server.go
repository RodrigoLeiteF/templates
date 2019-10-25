package main

import (
	"context"
	"log"
	"time"

	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	cloudkms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/pubsub"
	"github.com/getsentry/sentry-go"
	"github.com/lucsky/cuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	DB     *mongo.Database
	PubSub *pubsub.Client
	KMS    *cloudkms.KeyManagementClient
}

// RegisterServer registers the {{name}} API Server.
func RegisterServer(gs *grpc.Server, db *mongo.Database, ps *pubsub.Client, kms *cloudkms.KeyManagementClient) {
	server := &server{
		DB:     db,
		PubSub: ps,
		KMS:    kms,
	}
	{{name}}pb.Register{{name}}APIServer(gs, server)
}

{{#each functions}}
{{> function rpc=rpc service_name=../name request=request response=response}}
{{/each}}
