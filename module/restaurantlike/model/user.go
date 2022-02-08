package restaurantlikemodel

import (
	"food-delivery/common"
	"time"
)

type User struct {
	common.SimpleUser `json:",inline"`
	LikedAt           *time.Time `json:"created_at,omitempty" gorm:"created_at"`
}
