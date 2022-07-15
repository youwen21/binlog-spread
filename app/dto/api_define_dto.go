package dto

import (
	"binlog_spread/app/models"
)

type ApiDefineForm struct {
	models.ApiDefine
	PageParam
}

type ApiDefineResult struct {
	Count int64              `json:"count"`
	List  []models.ApiDefine `json:"list"`
}

type ApiDefineText struct {
	models.ApiDefine
}
