package models

type StateAbnormal struct {
	StateAbnormalId int       `gorm:"primary_key;" json:"state_abnormal_id"`
	EventType       int       `json:"event_type"`
	DbName          string    `json:"db_name"`
	TableName       string    `json:"table_name"`
	FieldName       string    `json:"field_name"`
	StateFrom       string    `json:"state_from"`
	StateTo         string    `json:"state_to"`
	IsDeleted       int       `json:"is_deleted"`
	CreatedAt       DateTime `json:"created_at" gorm:"autoUpdateTime;column:created_at;default:null"`
	UpdatedAt       DateTime `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;default:null" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	DeletedAt       DateTime `json:"deleted_at" xorm:"DATETIME"`
}
