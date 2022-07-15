package dto

import (
	"binlog_spread/app/models"
)

type StateForm struct {
	models.State
	PageParam
}

type StateResult struct {
	Count int64          `json:"count"`
	List  []models.State `json:"list"`
}

type StateText struct {
	models.State
	DbName         string `form:"db_name" json:"db_name" xorm:"default '' VARCHAR(255)"`
	TableName      string `form:"table_name" json:"table_name" xorm:"default '' VARCHAR(255)"`
	FieldName      string `form:"field_name" json:"field_name" xorm:"default '' VARCHAR(255)"`
	StateClassName string `form:"state_class_name" json:"state_class_name"`
}

type StateClassForm struct {
	models.StateClass
	PageParam
}

type StateClassQueryResult struct {
	Count int64               `json:"count"`
	List  []models.StateClass `json:"list"`
}

type StateDirectionForm struct {
	models.StateDirection
	PageParam
}

type StateDirectionQueryResult struct {
	Count int64                   `json:"count"`
	List  []models.StateDirection `json:"list"`
}

type StateDirectionText struct {
	models.StateDirection
	StateFromDesc string `form:"state_from_desc" json:"state_from_desc"`
	StateToDesc   string `form:"state_to_desc" json:"state_to_desc"`
}

type StateAbnormalForm struct {
	models.StateAbnormal
	PageParam
}

type StateAbnormalQueryResult struct {
	Count int64                  `json:"count"`
	List  []models.StateAbnormal `json:"list"`
}
