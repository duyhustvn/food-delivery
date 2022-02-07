package restaurantmodel

import "errors"

var (
	ErrRestaurantNameCannotBeBlank = errors.New("name cannot be blank")
)
