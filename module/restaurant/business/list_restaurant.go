package bizrestaurant

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type GetRestaurantLikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantBiz struct {
	store     ListRestaurantStore
	likeStore GetRestaurantLikeStore
}

func NewListRestaurantBiz(store ListRestaurantStore, likeStore GetRestaurantLikeStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store, likeStore: likeStore}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, "User", "Test")

	if err != nil {
		return nil, err
	}

	ids := make([]int, len(result))
	for i, item := range result {
		ids[i] = item.ID
	}

	if likeCounts, err := biz.likeStore.GetRestaurantLikes(ctx, ids); err == nil {
		for i, item := range result {
			result[i].LikeCount = likeCounts[item.ID]
		}
	}

	return result, nil
}
