package models

type ApiDefine struct {
	ApiDefineId int      `gorm:"primaryKey;primary_key;" json:"api_define_id" form:"api_define_id"`
	ApiType     string   `json:"api_type" form:"api_type"`
	ApiTag      string   `json:"api_tag" form:"api_tag"`
	ApiName     string   `json:"api_name" form:"api_name" binding:"required" validate:"required,min=2"`
	StreamIds   string   `json:"stream_ids" form:"stream_ids"`
	ApiVersion  string   `json:"api_version" form:"api_version"`
	ApiLink     string   `json:"api_link" form:"api_link"`
	Comment     string   `json:"comment" form:"comment"`
	IsDeleted   int      `json:"is_deleted" form:"is_deleted"`
	DeletedAt   DateTime `gorm:"autoCreateTime" json:"deleted_at" form:"deleted_at"`
	CreatedAt   DateTime `gorm:"autoUpdateTime;column:created_at;default:null" json:"created_at" form:"created_at"`
	UpdatedAt   DateTime `gorm:"autoUpdateTime;column:updated_at;default:null" json:"updated_at" form:"updated_at"`
}
