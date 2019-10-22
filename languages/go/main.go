package main

import (
	"context"
	"fmt"
	"log"
	"net"

	{{name}}pb "github.com/upnid/protobuf/gen/go/upnid/{{name}}/v1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/caarlos0/env/v6"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port        string `env:"SERVER_PORT" envDefault:"5080"`
	Host        string `env:"SERVER_HOST" envDefault:"localhost"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	SentryDsn   string `env:"SENTRY_DSN"`
	MongoURI    string `env:"MONGO_URI" envDefault:"mongodb://mongo.gateway.svc.cluster.local"`
}

func initDatabase(ctx context.Context, uri string) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return client
}

func initGrpcServer(server Server, port string, host string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	log.Print("gRPC server started on http://" + host + ":" + port)
	{{name}}pb.RegisterLogAPIServer(srv, &server)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		log.Fatalf("Could not parse config: %v", err)
	}

	server := Server{
		clients: Clients{
			Db: initDatabase(context.Background(), cfg.MongoURI).Database("upnid"),
		},
	}

	initGrpcServer(server, cfg.Port, cfg.Host)
}
