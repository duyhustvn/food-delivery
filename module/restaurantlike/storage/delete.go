package restaurantlikestorage

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Where("user_id = ? AND restaurant_id = ?", data.UserId, data.RestaurantId).Delete(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
