package ginrestaurant

import (
	"food-delivery/common"
	"food-delivery/component/appctx"
	bizrestaurant "food-delivery/module/restaurant/business"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		//go func() {
		//	defer common.Recover()
		//
		//	arr := []int{}
		//	log.Println(arr[0])
		//}()

		//id, err := strconv.Atoi(c.Param("id"))

		id, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())
		biz := bizrestaurant.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), int(id.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
