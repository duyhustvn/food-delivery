package restaurantmodel

import "food-delivery/common"

type RestaurantCreate struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	Address string         `json:"address" gorm:"column:addr;"`
	OwnerId int            `json:"owner_id" gorm:"column:owner_id;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
