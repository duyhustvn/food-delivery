package restaurantmodel

import (
	"food-delivery/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	OwnerId int            `json:"owner_id" gorm:"column:owner_id;"`
	Address string         `json:"address" gorm:"column:addr;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (data *Restaurant) Mask(isOwnerOrAdmin bool) {
	data.SQLModel.Mask(common.DbTypeRestaurant)
}
