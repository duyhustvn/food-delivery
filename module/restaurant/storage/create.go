package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
