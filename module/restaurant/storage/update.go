package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"

	"gorm.io/gorm"
)

// id: restaurant id
func (s *sqlStore) Update(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

// id: restaurant id
func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

// id: restaurant id
func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
