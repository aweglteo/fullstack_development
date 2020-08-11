package grpcc

import (
	"log"
	"os"

	pb "github.com/aweglteo/fullstack_development/api/external/scraping"
	"google.golang.org/grpc"
)

var GRPCClient *pb.ScrapingClient
var GRPConn *grpc.ClientConn

func Connect() {
	GRPC_SERVER_HOST := os.Getenv("PARSER_HOST") + ":" + os.Getenv("PARSER_PORT")
	conn, err := grpc.Dial(GRPC_SERVER_HOST, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	client := pb.NewScrapingClient(conn)

	GRPConn = conn
	GRPCClient = &client
}

func CloseConn() {
	GRPConn.Close()
}
