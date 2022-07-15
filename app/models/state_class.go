package models

type StateClass struct {
	StateClassId  int      `gorm:"primary_key;" form:"state_class_id" json:"state_class_id" xorm:"not null pk autoincr INT(11)"`
	DbName        string   `form:"db_name" json:"db_name" xorm:"default '' VARCHAR(255)"`
	TableName     string   `form:"table_name" json:"table_name" xorm:"default '' VARCHAR(255)"`
	FieldName     string   `form:"field_name" json:"field_name" xorm:"default '' VARCHAR(255)"`
	StateDescribe string   `form:"state_describe" json:"state_describe" xorm:"default '' VARCHAR(255)"`
	StateName     string   `form:"state_name" json:"state_name" xorm:"default '' VARCHAR(255)"`
	CreatedAt     DateTime `form:"created_at" json:"created_at" xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdatedAt     DateTime `form:"updated_at" json:"updated_at" xorm:"DATETIME"`
	IsDeleted     int      `form:"is_deleted" json:"is_deleted" xorm:"default 0 TINYINT(4)"`
	DeletedAt     DateTime `form:"deleted_at" json:"deleted_at" xorm:"DATETIME"`
}
