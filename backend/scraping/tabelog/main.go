package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// this module is scraping server via gRPC
func main() {
	port, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()

	// sampleUrl := "https://tabelog.com/tokyo/A1310/A131002/13181683/"
	// scrapingService := ScrapingService{}

	reflection.Register(server)

	log.Println("server start ... ")

	if err := server.Serve(port); err != nil {
		log.Fatalf("failed serve: %v", err)
	}
}
