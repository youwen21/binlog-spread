package models

type State struct {
	StateId        int      `gorm:"primary_key;" form:"state_id" json:"state_id"`
	StateClassId   int      `form:"state_class_id" json:"state_class_id"`
	StateValue     string   `form:"state_value" json:"state_value"`
	StateValueDesc string   `form:"state_value_desc" json:"state_value_desc"`
	IsDeleted      int      `form:"is_deleted" json:"is_deleted"`
	CreatedAt      DateTime `form:"created_at" json:"created_at"`
	UpdatedAt      DateTime `form:"updated_at" json:"updated_at"`
	DeletedAt      DateTime `form:"deleted_at" json:"deleted_at"`
}
