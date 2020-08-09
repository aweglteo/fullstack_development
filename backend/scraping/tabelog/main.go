package main

import (
	"log"
	"net"

	pb "github.com/aweglteo/fullstack_development/api/scraping"
	"github.com/aweglteo/fullstack_development/backend/scraping/tabelog/scraping"
	"google.golang.org/grpc"
)

// this module is scraping server via gRPC
func main() {
	port, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	ss := scraping.Server{}

	server := grpc.NewServer()
	pb.RegisterScrapingServer(server, &ss)

	log.Println("server start ... ")

	if err := server.Serve(port); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
