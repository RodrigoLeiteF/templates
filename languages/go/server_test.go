package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	{{name}} "github.com/upnid/protobuf/gen/go/upnid/{{user}}/v1"
)

var mgo *mongo.Database

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mgoclient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Could not connect to mongo: %v", err)
	}
	mgo = mgoclient.Database("upnid")

	//While use the same bd url
	err = mgo.Collection("{{name}}s").Drop(ctx)
}

func serverInstance(mgo *mongo.Database) server {
	return server{
		DB:     mgo,
	}
}

func TestCreate{{name}}Success(t *testing.T) {
}
