package gateway

import (
	"context"
	"fmt"

	pb "github.com/aweglteo/fullstack_development/api/external/scraping"
	"google.golang.org/grpc"

	"github.com/aweglteo/fullstack_development/api/domain"
	"github.com/jinzhu/gorm"
)

type RestaurantRepository struct {
	Conn     *gorm.DB
	GRPCConn *grpc.ClientConn
}

type Restaurant struct {
	gorm.Model
	Name string `gorm:not null`
}

func (rr *RestaurantRepository) Store(r domain.Restaurant) (id int, err error) {
	return 0, nil
}

func (rr *RestaurantRepository) FindAll() ([]domain.Restaurant, error) {
	return []domain.Restaurant{}, nil
}

func (rr *RestaurantRepository) FindByUser(user domain.User) ([]domain.Restaurant, error) {
	return []domain.Restaurant{}, nil
}

func (rr *RestaurantRepository) GetInfo(targetURL string) (domain.Restaurant, error) {

	fmt.Println("trargetLink:   ", targetURL)
	message := pb.GetScrapingRequest{TargetLink: targetURL}
	client := pb.NewScrapingClient(rr.GRPCConn)
	res, err := client.GetShopInfo(context.TODO(), &message)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)

	if err != nil {
		return domain.Restaurant{}, err
	}

	restaurant := domain.Restaurant{
		Name: res.Name,
	}

	return restaurant, nil
}
