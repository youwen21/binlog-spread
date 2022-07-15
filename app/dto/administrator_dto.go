package dto

import (
	"binlog_spread/app/models"
)

type AdminForm struct {
	models.Menu
	PageForm
}

type AdminResult struct {
	Count int64                   `json:"count"`
	List  []*models.Administrator `json:"list"`
}

type AdminText struct {
	models.Administrator
}
