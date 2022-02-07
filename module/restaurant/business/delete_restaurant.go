package bizrestaurant

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteStore interface {
	GetDataWithCondition(ctx context.Context, cond map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, updateData *restaurantmodel.RestaurantUpdate) error
}

type deleteRestaurantBiz struct {
	store DeleteStore
}

func NewDeleteRestaurantBiz(store DeleteStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int, isSoft bool) error {
	oldData, err := biz.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("restaurant has been deleted")
	}

	if isSoft {
		zero := 0
		if err := biz.store.Update(ctx, id, &restaurantmodel.RestaurantUpdate{Status: &zero}); err != nil {
			return err
		}
		return nil
	}

	// HARD DELETE
	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
