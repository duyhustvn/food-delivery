package reporestaurant

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

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore GetRestaurantLikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore, likeStore GetRestaurantLikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, likeStore: likeStore}
}

func (biz *listRestaurantRepo) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, moreKeys...)

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
