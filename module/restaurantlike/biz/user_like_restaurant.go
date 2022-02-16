package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
}

type IncreaseRestaurantCounterLike interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store     UserLikeRestaurantStore
	likeStore IncreaseRestaurantCounterLike
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, likeStore IncreaseRestaurantCounterLike) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurantlikemodel.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}
	go func() {
		defer common.Recover()
		_ = biz.likeStore.IncreaseLikeCount(ctx, data.RestaurantId)
	}()
	return nil
}
