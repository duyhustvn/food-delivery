package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	bizrestaurant "food-delivery/module/restaurant/business"
	restaurantmodel "food-delivery/module/restaurant/model"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&newRestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		newRestaurant.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())
		biz := bizrestaurant.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newRestaurant.ID))
	}
}
