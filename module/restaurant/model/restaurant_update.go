package restaurantmodel

import "food-delivery/common"

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:addr;"`
	Status  *int           `json:"-" gorm:"column:status;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantUpdate) Validate() error {
	if v := data.Name; v != nil && *v == "" {
		return ErrRestaurantNameCannotBeBlank
	}

	return nil
}
