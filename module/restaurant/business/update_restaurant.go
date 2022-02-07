package bizrestaurant

import (
	"context"
	"errors"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type UpdateStore interface {
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateStore
}

func NewUpdateRestaurantBiz(store UpdateStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := data.Validate(); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.EntityName, err)
	}

	oldData, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("restaurant has been deleted")
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
