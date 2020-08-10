// sample client of gRPC

package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aweglteo/fullstack_development/api/external/scraping"
	"google.golang.org/grpc"
)

func main() {

	postUrl := "https://tabelog.com/tokyo/A1310/A131002/13181683/"

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := pb.NewScrapingClient(conn)
	message := pb.GetScrapingRequest{TargetLink: postUrl}
	res, err := client.GetShopInfo(context.TODO(), &message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
