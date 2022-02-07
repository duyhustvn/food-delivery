package restaurantmodel

type Filter struct {
	CategoryId int `json:"category_id" form:"category_id"`
	OwnerId    int `json:"owner_id" form:"owner_id"`
}
