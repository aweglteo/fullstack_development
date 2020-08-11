package grpcc

import (
	"log"
	"os"

	"google.golang.org/grpc"
)

var GRPConn *grpc.ClientConn

func Connect() {
	GRPC_SERVER_HOST := os.Getenv("PARSER_HOST") + ":" + os.Getenv("PARSER_PORT")
	conn, err := grpc.Dial(GRPC_SERVER_HOST, grpc.WithInsecure())

	if err != nil {
		log.Fatal("client connection error:", err)
	}

	GRPConn = conn

}

func CloseConn() {
	GRPConn.Close()
}
