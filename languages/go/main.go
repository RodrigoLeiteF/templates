package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port        string `env:"SERVER_PORT"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	SentryDsn   string `env:"SENTRY_DSN"`
	MongoURI    string `env:"MONGO_URI" envDefault:"mongodb://mongo.gateway.svc.cluster.local"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		log.Fatalf("Could not parse config: %v", err)
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         cfg.SentryDsn,
		Environment: cfg.Environment,
	})

	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Could not listen to port: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mgoclient, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Could not connect to mongo: %v", err)
	}

	gs := grpc.NewServer()
	RegisterServer(gs, mgoclient.Database("upnid"))
	reflection.Register(gs)

	log.Print("Serving gRPC on http://localhost:" + os.Getenv("SERVER_PORT"))

	if e := gs.Serve(listener); e != nil {
		sentry.CaptureException(err)
		log.Fatalf("Could not serve grpc server: %v", err)
	}
}
