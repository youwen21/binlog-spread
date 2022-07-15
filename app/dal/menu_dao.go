package dal

import (
	"binlog_spread/app/dto"
	"binlog_spread/app/models"
	"binlog_spread/comps"
	"gorm.io/gorm"
	"time"
)

type menuDAO struct{
}

var (
	MenuDAO = &menuDAO{}

	//queryErr = errors.New("查询失败")
)

func (d *menuDAO) Query(f *dto.MenuForm) (*dto.MenuResult, error) {
	session := d.newSession()

	var total int64
	var list []models.Menu

	if f.IsDeleted != -1 {
		session.Where("is_deleted = ?", f.IsDeleted)
	}

	if err := session.Count(&total).Error; err != nil {
		return nil, err
	}

	if len(f.OrderBy) > 0 {
		for _, v := range f.OrderBy {
			session.Order(v)
		}
	}

	if err := session.Limit(f.Limit()).Offset(f.Offset()).Find(&list).Error; err != nil {
		return nil, err
	}

	return &dto.MenuResult{Count: total, List: list}, nil
}

func (d *menuDAO) Get(id int) (*models.Menu, error) {
	info := &models.Menu{}
	session := d.newSession()
	if err := session.Where(" menu_id= ?", id).First(info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (d *menuDAO) GetList(f *dto.MenuForm) ([]*models.Menu, error) {
	var results []*models.Menu

	session := d.newSession()
	if f.IsDeleted != -1 {
		session.Where("is_deleted= ?", f.IsDeleted)
	}
	err := session.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *menuDAO) Insert(menu *models.Menu) error {
	session := d.newSession()
	err := session.Create(&menu).Error
	return err
}

func (d *menuDAO) Update(menu *models.Menu) error {
	session := d.newSession()
	err := session.Where("menu_id = ?", menu.MenuId).UpdateColumns(menu).Error
	return err
}

func (d *menuDAO) Delete(id int) error {
	session := d.newSession()
	err := session.Where("menu_id = ?", id).Updates(map[string]interface{}{"is_deleted": 1, "deleted_at": time.Now()}).Error
	return err
}

func (d *menuDAO) UpdateByMap(id int, updated map[string]interface{}) (int64, error) {
	session := d.newSession()
	if err := session.Where("menu_id = ?", id).UpdateColumns(updated).Error; err != nil {
		return 0, err
	}
	return session.RowsAffected, nil
}

//func (d *menuDAO) BatchInsert(examineList []*models.Menu) (int64, error) {
//	return 0, nil
//}
//
//func (d *menuDAO) BatchUpdateByMap(idList []int, updated map[string]interface{}) error {
//	return nil
//}
//
//func (d *menuDAO) BatchDelete(idList []int) (int64, error) {
//	return 0, nil
//}

func (d *menuDAO) newEngine() *gorm.DB {
	return comps.GetDb()
}

func (d *menuDAO) newSession() *gorm.DB {
	return comps.GetSession().Table("menu")
}

