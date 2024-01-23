package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	db "pharmago/account/db/sqlc"
	"pharmago/account/gapi"
	"pharmago/utils"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kothar/go-backblaze"
	"github.com/longIdt2502/pharmago-microservice/account/pb"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

type server struct {
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	conn, err := sql.Open(config.DB_DRIVER_ACCOUNT, config.DB_SOURCE_ACCOUNT)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot open database")
	}

	utils.RunDBMigration(config.MigrationURL, config.DB_SOURCE_ACCOUNT)

	store := db.NewStore(conn)

	run
}

func runServerGRPC(config utils.Config, store *db.StoreAccout, taskDistributor woker.TaskDistributor, b2Bucket *backblaze.Bucket) {
	server, err := gapi.NewServer(config, store, taskDistributor, b2Bucket)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(config2.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterPharmagoServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start GRPC server")
	}
}

func runGatewayServer(config utils.Config, store *db.StoreAccout, taskDistributor woker.TaskDistributor, b2Bucket *backblaze.Bucket) {
	server, err := gapi.NewServer(config, store, taskDistributor, b2Bucket)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	option := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(option)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterPharmagoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register handle server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	//fs := http.FileServer(http.Dir("./doc/swagger"))
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}
	swaggerHandle := http.StripPrefix("/swagger", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandle)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := config2.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server")
	}
}
