package service

import (
	"binlog_spread/app/dal"
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
)

type menuSrv struct{}

var (
	MenuSrv = &menuSrv{}
)

func (srv *menuSrv) Query(f *dto.MenuForm) (*dto.MenuResult, error) {
	return dal.MenuDAO.Query(f)
}

func (srv *menuSrv) Get(id int) (*models.Menu, error) {
	return dal.MenuDAO.Get(id)
}

func (srv *menuSrv) GetMenuTree() ([]dto.MenuTreeItem, error) {
	form := &dto.MenuForm{
		Menu: models.Menu{
			IsDeleted: 0,
		},
		PageParam: dto.PageParam{
			PageNumber:  1,
			PageSize: 1000,
		},
	}
	list, err := dal.MenuDAO.GetList(form)
	if err != nil {
		return nil, err
	}

	menuList := dto.MenuList(list)
	tree := menuList.ProcessToTree(0, 0)
	return tree, nil
}
