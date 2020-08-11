package grpcc

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

var GrpcConn *grpc.ClientConn

// MEMO: Don't forget to close connection
func Connect() {
	GRPC_SERVER_HOST := os.Getenv("PARSER_HOST") + ":" + os.Getenv("PARSER_PORT")
	GrpcConn, err := grpc.Dial(GRPC_SERVER_HOST, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
}
