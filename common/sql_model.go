package common

import "time"

type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id;"`
	FakeId    UID        `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
}

func (sm *SQLModel) GenUID(objType int) {
	sm.FakeId = NewUID(uint32(sm.ID), objType, 1)
}
