package restaurantlikebiz

import (
	"context"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := biz.store.Delete(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}
	return nil
}
