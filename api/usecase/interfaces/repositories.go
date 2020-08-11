package interfaces

import "github.com/aweglteo/fullstack_development/api/domain"

type RestaurantRepository interface {
	Store(domain.Restaurant) (id int, err error)
	FindByUser(domain.User) ([]domain.Restaurant, error)
	FindAll() ([]domain.Restaurant, error)
}
