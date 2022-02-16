package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *restaurantlikemodel.Like) error
}

type DecreaseRestaurantCounterLikeStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store     UserDislikeRestaurantStore
	likeStore DecreaseRestaurantCounterLikeStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, likeStore DecreaseRestaurantCounterLikeStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := biz.store.Delete(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotDislikeRestaurant(err)
	}

	go func() {
		defer common.Recover()
		_ = biz.likeStore.DecreaseLikeCount(ctx, data.RestaurantId)
	}()

	return nil
}
