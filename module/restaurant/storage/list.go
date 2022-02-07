package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	var result []restaurantmodel.Restaurant

	db = db.Where("status in (?)", 1)

	if filter.OwnerId > 0 {
		db = db.Where("owner_id = ?", filter.OwnerId)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		if moreKeys[i] == "User" {
			db = db.Preload("User")
		}
	}

	if err := db.
		Limit(paging.Limit).
		Offset((paging.Page - 1) * paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
