package usecase

import (
	"github.com/aweglteo/fullstack_development/api/domain"
	"github.com/aweglteo/fullstack_development/api/usecase/interfaces"
)

type RestaurantInteractor struct {
	RestaurantRepository interfaces.RestaurantRepository
}

func (i *RestaurantInteractor) Add(r domain.Restaurant) {
	i.RestaurantRepository.Store(r)
}

func (i *RestaurantInteractor) GetInfo(targetURL string) (domain.Restaurant, error) {
	return i.RestaurantRepository.GetInfo(targetURL)
}
