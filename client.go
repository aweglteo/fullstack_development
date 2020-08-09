// sample client of gRPC

package main

import (
	pb "github.com/aweglteo/fullstack_development/api/scraping"
	"google.golang.org/grpc"
)

func main() {
	conn, err: = grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()


}