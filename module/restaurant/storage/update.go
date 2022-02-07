package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}
