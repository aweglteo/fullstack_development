package gateway

import (
	pb "github.com/aweglteo/fullstack_development/api/external/scraping"

	"github.com/aweglteo/fullstack_development/api/domain"
	"github.com/jinzhu/gorm"
)

type RestaurantRepository struct {
	Conn       *gorm.DB
	GRPCClient *pb.ScrapingClient
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
	return domain.Restaurant{}, nil
}
