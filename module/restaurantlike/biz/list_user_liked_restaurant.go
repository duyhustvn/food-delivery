package restaurantlikebiz

import (
	"context"
	"food-delivery/common"
	restaurantlikemodel "food-delivery/module/restaurantlike/model"
)

type ListUserLikeRestaurantStore interface {
	ListUserLiked(ctx context.Context, conditions map[string]interface{}, filter *restaurantlikemodel.Filter, paging *common.Paging) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUsers(ctx context.Context, filter *restaurantlikemodel.Filter, paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := biz.store.ListUserLiked(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
